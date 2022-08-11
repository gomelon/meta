package testdata

import (
	"context"
	"time"
)

//SomeStruct
//aop:iface
type SomeStruct struct {
}

func (s *SomeStruct) PublicMethod(ctx context.Context, id int64) (string, error) {
	return "nil", nil
}

func (s *SomeStruct) privateMethod(ctx context.Context, time time.Time) (int32, error) {
	return 0, nil
}

//NoneMethodStruct
//aop:iface
type NoneMethodStruct struct {
}
