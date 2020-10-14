package gribbon

import (
	"errors"
	"strings"
)

var (
	errInvalidCap  = errors.New("cap is invalid")            // 容量异常
	errWrongSubmit =errors.New("submit function is illegal") // 提交异常方法
	errOverMaxCap  = errors.New("link is over max cap")      // 到达最大容量限制
	errClosed      = errors.New("link is closed")            // 已经关闭
)

func IsGribbonErr(err error)bool{
	return IsInvalidCap(err)||IsWrongSubmit(err)||IsOverMaxCap(err)||IsClosed(err)
}

func IsInvalidCap(err error)bool{
	return strings.EqualFold(err.Error(),errInvalidCap.Error())
}

func IsWrongSubmit(err error)bool{
	return strings.EqualFold(err.Error(),errWrongSubmit.Error())
}

func IsOverMaxCap(err error) bool {
	return strings.EqualFold(err.Error(), errOverMaxCap.Error())
}

func IsClosed(err error) bool {
	return strings.EqualFold(err.Error(), errClosed.Error())
}
