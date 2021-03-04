package scheme

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type Value struct {
	Float   *float64 `  @Float`
	Bool    *Boolean `| @("true" | "false")`
	AString *string  `| @String`
	Func    SchemeFunction
}

func (v Value) ToFloat() (float64, error) {
	if v.Float != nil {
		return *v.Float, nil
	}

	return 0, fmt.Errorf("not a number (%s)", v.String())
}

func (v Value) Primitive() interface{} {
	if v.Float != nil {
		return *v.Float
	}

	if v.Bool != nil {
		return *v.Bool
	}

	if v.AString != nil {
		s := *v.AString
		return s
	}

	panic("help")
}

func (v Value) String() string {
	if v.Float != nil {
		return fmt.Sprintf("%v", *v.Float)
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
	Call   []*Expression `"(" @@+ ")"|`
	Val    *Value        `@@|`
	Var    *string       `@Ident|`
	Symbol *string       `@Symbol`
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
	schemeLexer = stateful.MustSimple([]stateful.Rule{
		{"Ident", `[a-zA-Z]\w*`, nil},
		{"Float", `[-+]?\d*\.?\d+([eE][-+]?\d+)?`, nil},
		{"String", `"(\\"|[^"])*"`, nil},
		{"Whitespace", `[ \t\n\r]+`, nil},
		{"Symbol", "(=|<=|\\+|\\*|-|/)", nil},
		{"EOL", `[\n\r]+`, nil},
		{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	})

	schemeParser = participle.MustBuild(
		&Expression{},
		participle.Lexer(schemeLexer),
		participle.Unquote("String"),
		participle.Elide("Whitespace"),
	)
)

func Parse(s string) (*Expression, error) {

	e := &Expression{}
	err := schemeParser.ParseString("", s, e)
	return e, err
}
