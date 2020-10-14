package gribbon

import "context"

type node struct {
	data *worker
	next *node
}

func newNode(d *worker, n *node) *node {
	return &node{data: d, next: n}
}

func (n *node) insertTail(nd *node) {
	n.next = nd
}

func (n *node) isWorking() bool {
	return n.data.isWorking()
}

func (n *node) init(c context.Context, a interface{}) {
	n.data.init(c, a)
}

func (n *node) submit(f func(context.Context)) {
	n.data.submit(f)
}

func (n *node) submitWithArg(f func(context.Context, interface{})) {
	n.data.submitWithArg(f)
}

func (n *node) run() {
	n.data.run()
}

func (n *node) runWithArg() {
	n.data.runWithArg()
}

func (n *node) close() {
	n.data.close()
}
