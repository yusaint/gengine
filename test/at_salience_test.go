package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/yusaint/gengine/builder"
	"github.com/yusaint/gengine/context"
	"github.com/yusaint/gengine/engine"
)

func Test_at_salience(t *testing.T) {

	dataContext := context.NewDataContext()
	dataContext.Add("println", fmt.Println)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(`
rule "1" salience 10
begin
	println(@sal)
end
rule "2" 
begin
	println(@sal)
end
`)

	if err != nil {
		panic(err)
	}
	eng := engine.NewGengine()

	start := time.Now().UnixNano()
	// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
	err = eng.Execute(ruleBuilder, true)
	end := time.Now().UnixNano()
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("execute rule cost %d ns", end-start))

}
