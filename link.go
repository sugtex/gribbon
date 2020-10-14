package gribbon

import (
	"context"
	"sync"
)

type goLink struct {
	cap       uint8
	len       uint8
	isClose   bool
	isHaveArg bool
	lock      sync.Mutex
	head      *node
}

func NewGoLink(c uint8, isHaveArg bool) (*goLink, error) {
	if c <= 0 {
		return nil, errInvalidCap
	}
	return &goLink{
		cap:       c,
		len:       0,
		isHaveArg: isHaveArg,
		isClose:   false,
		head:      newNode(nil, nil),
	}, nil
}

func (g *goLink) Len() uint8 {
	return g.len
}

func (g *goLink) Submit(ctx context.Context, f func(context.Context)) error {
	if g.isArgLink() {
		return errWrongSubmit
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.isClose {
		return errClosed
	}

	isFind, node := g.tryForeachIdleWorker()
	if isFind {
		node.init(ctx, nil)
		node.submit(f)
		return nil
	}

	newNode := newNode(newWorker(ctx, nil), nil)
	if err := g.tryCreateNew(node, newNode, newNode.run); err != nil {
		return err
	}
	newNode.submit(f)
	return nil
}

func (g *goLink) SubmitWithArg(ctx context.Context, arg interface{}, f func(context.Context, interface{})) error {
	if !g.isArgLink() {
		return errWrongSubmit
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.isClose {
		return errClosed
	}

	isFind, node := g.tryForeachIdleWorker()
	if isFind {
		node.init(ctx, arg)
		node.submitWithArg(f)
		return nil
	}

	newNode := newNode(newWorker(ctx, arg), nil)
	if err := g.tryCreateNew(node, newNode, newNode.runWithArg); err != nil {
		return err
	}
	newNode.submitWithArg(f)
	return nil
}

func (g *goLink) isArgLink() bool {
	return g.isHaveArg
}

func (g *goLink) Close() error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.isClose {
		return errClosed
	}
	g.isClose = true
	go g.asyncCloseLink()
	return nil
}

func (g *goLink) linkHead() *node {
	return g.head
}

// 尝试寻找空闲的node
func (g *goLink) tryForeachIdleWorker() (bool, *node) {
	node := g.linkHead().next
	pre := g.linkHead()
	for node != nil {
		if !node.isWorking() {
			return true, node
		}
		node = node.next
		if node != nil {
			pre = node
		}
	}
	return false, pre
}

// 尝试构建node并返回
func (g *goLink) tryCreateNew(pre *node, new *node, run func()) error {
	if g.isOverMaxCap() {
		return errOverMaxCap
	}
	pre.insertTail(new)
	go run()
	g.len++
	return nil
}

// 是否达到最大容量限制
func (g *goLink) isOverMaxCap() bool {
	return g.len >= g.cap
}

// 异步处理关闭链路
func (g *goLink) asyncCloseLink() {
	node := g.linkHead().next
	for node != nil {
		node.close()
		node = node.next
	}
}
