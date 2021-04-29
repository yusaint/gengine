package base

import (
	"errors"
	"reflect"

	"github.com/yusaint/gengine/context"
)

type ElseStmt struct {
	StatementList *Statements
}

func (e *ElseStmt) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error, bool) {

	if e.StatementList != nil {
		return e.StatementList.Evaluate(dc, Vars)
	} else {
		return reflect.ValueOf(nil), nil, false
	}
}

func (e *ElseStmt) AcceptStatements(stmts *Statements) error {
	if e.StatementList == nil {
		e.StatementList = stmts
		return nil
	}
	return errors.New("ElseStmt set twice! ")
}
