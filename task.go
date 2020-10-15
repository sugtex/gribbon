package gribbon

import "context"

type task struct {
	ctx context.Context
	arg interface{}

	f   func(context.Context)
	fwa func(context.Context, interface{})
}

func newTask(c context.Context, a interface{}, f func(context.Context), fwa func(context.Context, interface{})) *task {
	return &task{
		ctx: c,
		arg: a,
		f:   f,
		fwa: fwa,
	}
}

func (t *task) run() {
	t.f(t.ctx)
}

func (t *task) runWithArg() {
	t.fwa(t.ctx, t.arg)
}