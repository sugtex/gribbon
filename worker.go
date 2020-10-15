package gribbon

import (
	"log"
)

type worker struct {
	isBusy bool
	taskC  chan *task
}

func newWorker() *worker {
	return &worker{
		isBusy: false,
		taskC:  make(chan *task),
	}
}

func (w *worker) isWorking() bool {
	return w.isBusy
}

func (w *worker) submit(t *task) {
	w.taskC <- t
}

func (w *worker) run() {
	defer w.recoverPanic()
	for t := range w.taskC {
		if t == nil {
			return
		}
		w.pre()
		t.run()
		w.after()
	}
}

func (w *worker) runWithArg() {
	defer w.recoverPanic()
	for t := range w.taskC {
		if t == nil {
			return
		}
		w.pre()
		t.runWithArg()
		w.after()
	}
}

func (w *worker) recoverPanic() {
	if err := recover(); err != nil {
		log.Fatalf("worker err:%s", err)
	}
}

func (w *worker) pre() {
	w.isBusy = true
}

func (w *worker) close() {
	w.taskC <- nil
}

func (w *worker) after() {
	w.isBusy = false
}