package test

import (
	"testing"

	"github.com/yusaint/gengine/builder"
	"github.com/yusaint/gengine/context"
	"github.com/yusaint/gengine/engine"
)

func Test_lexer(t *testing.T) {

	lexer_rule := `
rule "test" salience 1
begin
规则管理
end
`
	dataContext := context.NewDataContext()
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(lexer_rule)
	if e1 != nil {
		panic(e1)
	}

	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	if e != nil {
		panic(e)
	}

}
