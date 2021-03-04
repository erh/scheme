package scheme

import (
	"fmt"
	"strconv"

	"github.com/alecthomas/participle/v2"
)

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type Value struct {
	Float   *float64 `  @Float`
	Int     *int     `| @Int`
	Bool    *Boolean `| @("true" | "false")`
	AString *string  `| @String`
	Func    SchemeFunction
}

func (v Value) ToFloat() (float64, error) {
	if v.Float != nil {
		return *v.Float, nil
	}

	if v.Int != nil {
		return float64(*v.Int), nil
	}

	return 0, fmt.Errorf("not a number")
}

func (v Value) Primitive() interface{} {
	if v.Float != nil {
		return *v.Float
	}

	if v.Int != nil {
		return *v.Int
	}

	if v.Bool != nil {
		return *v.Bool
	}

	if v.AString != nil {
		s := *v.AString
		return s[1 : len(s)-1]
	}

	panic("help")
}

func (v Value) String() string {
	if v.Float != nil {
		return fmt.Sprintf("%v", *v.Float)
	}

	if v.Int != nil {
		return strconv.Itoa(*v.Int)
	}

	if v.AString != nil {
		return *v.AString
	}

	if v.Bool != nil {
		return fmt.Sprintf("%v", *v.Bool)
	}

	if v.Func != nil {
		return "func"
	}

	panic("what am i!")
}

type Expression struct {
	Call   []*Expression `"(" @@+ ")" |`
	Val    *Value        `@@ |`
	Var    *string       `@Ident |`
	Symbol *string       `@("=" | "<" "=" | ">" "=" | "<" | ">" | "!" | "=" | "+" | "-" | "*" | "/" )`
}

func (e Expression) String() string {
	if e.Val != nil {
		return e.Val.String()
	}

	if e.Var != nil {
		return *e.Var
	}

	if e.Symbol != nil {
		return *e.Symbol
	}

	if len(e.Call) > 0 {
		s := "("
		for _, x := range e.Call {
			s += x.String() + " "
		}
		s += ")"
		return s
	}

	panic("why am i alive")
}

var (
	schemeParser = participle.MustBuild(&Expression{})
)

func Parse(s string) (*Expression, error) {
	e := &Expression{}
	err := schemeParser.ParseString("", s, e)
	return e, err
}
