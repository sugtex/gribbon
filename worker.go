package gribbon

import (
	"context"
	"log"
)

type worker struct {
	isBusy bool
	ctx    context.Context

	task chan func(context.Context)

	arg         interface{}
	taskWithArg chan func(context.Context, interface{})
}

func newWorker(c context.Context, a interface{}) *worker {
	return &worker{
		isBusy:      false,
		ctx:         c,
		task:        make(chan func(context.Context)),
		arg:         a,
		taskWithArg: make(chan func(context.Context, interface{})),
	}
}

func (w *worker) init(c context.Context, a interface{}) {
	w.ctx, w.arg = c, a
}

func (w *worker) isWorking() bool {
	return w.isBusy
}

func (w *worker) submit(f func(context.Context)) {
	w.task <- f
}

func (w *worker) submitWithArg(f func(context.Context, interface{})) {
	w.taskWithArg <- f
}

func (w *worker) run() {
	defer w.recoverPanic()
	for t := range w.task {
		if t == nil {
			return
		}
		w.isBusy = true
		t(w.ctx)
		w.reset()
	}
}

func (w *worker) runWithArg() {
	defer w.recoverPanic()
	for t := range w.taskWithArg {
		if t == nil {
			return
		}
		w.isBusy = true
		t(w.ctx, w.arg)
		w.reset()
	}
}

func (w *worker) recoverPanic() {
	if err := recover(); err != nil {
		log.Fatalf("worker err:%s", err)
	}
}

func (w *worker) close() {
	w.task <- nil
}

func (w *worker) reset() {
	w.isBusy, w.ctx = false, nil
}
