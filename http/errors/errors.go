package errors

import (
	"net/http"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 401
func NotLoginError(msg ...string) Error {
	return _build(http.StatusUnauthorized, "unauthorized", msg...)
}

// 400
func BadRequestError(msg ...string) Error {
	return _build(http.StatusBadRequest, "bad request", msg...)
}

// 403
func NoPrivError(msg ...string) Error {
	return _build(http.StatusForbidden, "forbidden", msg...)
}

// 500
func InternalServerError(msg ...string) Error {
	return _build(http.StatusInternalServerError, "internal server error", msg...)
}

func _build(code int, defval string, custom ...string) Error {
	msg := defval
	if len(custom) > 0 {
		msg = custom[0]
	}
	return Error{
		Code: code,
		Msg:  msg,
	}
}

func MaybePanic(err error) {
	if err != nil {
		panic(Error{Msg: err.Error()})
	}
}

func Panic(msg string) {
	panic(Error{Msg: msg})
}
