package testdata

import (
	"time"
)

const IntConst = 1

var IntVar int

var StringVar string

type SimpleStruct struct {
	Name string
}

func (s SimpleStruct) Method() {
}

func (s SimpleStruct) MethodWithParamAndResult(name string) int {
	return 0
}

func (s SimpleStruct) MethodWithParamAndNameResult(name string) (r int) {
	return
}

func (s *SimpleStruct) PointerMethod() {
}

type SimpleInterface interface {
	Method()
	MethodWithParamAndResult(name string) int
	MethodWithParamAndNameResult(name string) (r int)
}

func SimpleFunc() {

}

var StructVar SimpleStruct

type OuterPackageVar time.Time

//CommentVar example
var CommentVar int64

var NoneCommentVar int64