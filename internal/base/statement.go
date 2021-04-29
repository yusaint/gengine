package base

import (
	"errors"
	"reflect"

	"github.com/yusaint/gengine/context"
)

type Statement struct {
	DeferStmt      *DeferStmt
	IfStmt         *IfStmt
	MethodCall     *MethodCall
	FunctionCall   *FunctionCall
	ThreeLevelCall *ThreeLevelCall
	Assignment     *Assignment
	ConcStatement  *ConcStatement
}

func (s *Statement) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error, bool) {

	if s.DeferStmt != nil {
		defer s.DeferStmt.Evaluate(dc, Vars)
		return reflect.ValueOf(nil), nil, false
	}

	if s.IfStmt != nil {
		return s.IfStmt.Evaluate(dc, Vars)
	}

	if s.MethodCall != nil {
		v, e := s.MethodCall.Evaluate(dc, Vars)
		return v, e, false
	}

	if s.FunctionCall != nil {
		v, e := s.FunctionCall.Evaluate(dc, Vars)
		return v, e, false
	}

	if s.Assignment != nil {
		v, e := s.Assignment.Evaluate(dc, Vars)
		return v, e, false
	}

	if s.ConcStatement != nil {
		v, e := s.ConcStatement.Evaluate(dc, Vars)
		return v, e, false
	}

	if s.ThreeLevelCall != nil {
		v, e := s.ThreeLevelCall.Evaluate(dc, Vars)
		return v, e, false
	}

	return reflect.ValueOf(nil), errors.New("Statement evaluate error!"), false
}

func (s *Statement) AcceptFunctionCall(funcCall *FunctionCall) error {
	if s.FunctionCall == nil {
		s.FunctionCall = funcCall
		return nil
	}
	return errors.New("FunctionCall already defined! ")
}

func (s *Statement) AcceptMethodCall(methodCall *MethodCall) error {
	if s.MethodCall == nil {
		s.MethodCall = methodCall
		return nil
	}
	return errors.New("MethodCall already defined! ")
}

func (s *Statement) AcceptThreeLevelCall(threeLevelCall *ThreeLevelCall) error {
	if s.ThreeLevelCall == nil {
		s.ThreeLevelCall = threeLevelCall
		return nil
	}
	return errors.New("threeLevelCall already defined! ")
}

func (s *Statement) AcceptAssignment(assignment *Assignment) error {
	if s.Assignment == nil {
		s.Assignment = assignment
		return nil
	}
	return errors.New("Assignment already defined! ")
}
