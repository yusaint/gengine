package base

import (
	"reflect"
	"github.com/bilibili/gengine/context"
)

type DeferStmt struct {
	MethodCall   *MethodCall
	FunctionCall *FunctionCall
}

func (s *DeferStmt) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error, bool) {
	if s.MethodCall != nil {
		v, e := s.MethodCall.Evaluate(dc, Vars)
		return v, e, false
	}

	if s.FunctionCall != nil {
		v, e := s.FunctionCall.Evaluate(dc, Vars)
		return v, e, false
	}
	return reflect.ValueOf(nil), nil, false
}

func (s *DeferStmt) AcceptFunctionCall(funcCall *FunctionCall) error {
	s.FunctionCall = funcCall
	return nil
}

func (s *DeferStmt) AcceptMethodCall(methodCall *MethodCall) error {
	s.MethodCall = methodCall
	return nil
}
