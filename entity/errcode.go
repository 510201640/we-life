package entity

import (
	"errors"
	"fmt"
)

type IErrCode interface {
	Code() func() int
	Msg() func() string
}

type ErrCode struct {
	Code int
	Msg  string
	Err  error
}

func NewErrCode(code int, msg string) *ErrCode {
	ec := new(ErrCode)
	ec.Code = code
	ec.Msg = msg
	return ec
}

func (e *ErrCode) ReplaceMsg(msg string) *ErrCode {
	return NewErrCode(e.Code, msg)
}

func (e *ErrCode) AddMsg(msg string) *ErrCode {
	ec := new(ErrCode)
	ec.Code = e.Code
	ec.Msg = fmt.Sprintf("%s:%s", e.Msg, msg)
	return ec
}
func (e *ErrCode) FromError(err error) *ErrCode {
	ec := new(ErrCode)
	ec.Code = e.Code
	ec.Msg = e.Msg
	ec.Err = err
	return ec
}

func (e *ErrCode) ToResult() *Result {
	return NewResult(e.Code, e.Msg, nil)
}

func (e *ErrCode) IsEmpty() bool {
	return e == nil
}

func (e *ErrCode) IsNotEmpty() bool {
	return e != nil
}

func (e *ErrCode) New() *ErrCode {
	errCode := &ErrCode{
		Code: e.Code,
		Msg:  e.Msg,
		Err:  errors.New(fmt.Sprintf("code(%d),msg(%s)", e.Code, e.Msg)),
	}
	return errCode
}
