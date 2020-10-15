package gribbon

import (
	"context"
	"math"
)

type iWorker interface {
	submit(*task)
	isWorking() bool
	run()
	runWithArg()
	close()
}

var (
	defaultLink, _       = NewGoLink(math.MaxUint8, false)
	defaultLinkWitArg, _ = NewGoLink(math.MaxUint8, true)
)

func Submit(ctx context.Context, f func(context.Context)) error {
	return defaultLink.Submit(ctx, f)
}

func SubmitWithArg(ctx context.Context, arg interface{}, f func(context.Context, interface{})) error {
	return defaultLinkWitArg.SubmitWithArg(ctx, arg, f)
}

func Len() uint8 {
	return defaultLink.Len()
}

func ArgLinkLen() uint8 {
	return defaultLinkWitArg.Len()
}
