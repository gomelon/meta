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

func (s SimpleStruct) StructMethod() {

}

type PointerSimpleStruct struct {
	Name string
}

func (s *PointerSimpleStruct) StructMethod() {

}

type SimpleInterface interface {
	InterfaceMethod()
}

func SimpleFunc() {

}

var StructVar SimpleStruct

type OuterPackageVar time.Time

//CommentVar example
var CommentVar int64

var NoneCommentVar int64
