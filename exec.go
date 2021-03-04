package scheme

import (
	"fmt"
)

type Scope map[string]*Value

type SchemeFunction func(args []*Value, scope Scope) (*Value, error)

var (
	symbols = map[string]*Value{}
)

func init() {
	symbols["+"] = &Value{Func: func(args []*Value, scope Scope) (*Value, error) {
		total := 0.0

		for _, a := range args {
			n, err := a.ToFloat()
			if err != nil {
				return nil, err
			}

			total += n
		}

		return &Value{Float: &total}, nil
	}}

	symbols["*"] = &Value{Func: func(args []*Value, scope Scope) (*Value, error) {
		total := 0.0

		for idx, a := range args {
			n, err := a.ToFloat()
			if err != nil {
				return nil, err
			}

			if idx == 0 {
				total = n
			} else {
				total *= n
			}
		}

		return &Value{Float: &total}, nil
	}}

	symbols["-"] = &Value{Func: func(args []*Value, scope Scope) (*Value, error) {
		total := 0.0

		for idx, a := range args {
			n, err := a.ToFloat()
			if err != nil {
				return nil, err
			}

			if idx == 0 {
				total = n
			} else {
				total = total - n
			}
		}

		return &Value{Float: &total}, nil
	}}

	symbols["/"] = &Value{Func: func(args []*Value, scope Scope) (*Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("/ requres exactly 2 args")
		}

		a, err := args[0].ToFloat()
		if err != nil {
			return nil, err
		}

		b, err := args[1].ToFloat()
		if err != nil {
			return nil, err
		}

		if b == 0 {
			return nil, fmt.Errorf("cannot divide by 0")
		}

		res := a / b

		return &Value{Float: &res}, nil
	}}

}

func evalAll(es []*Expression, s Scope) ([]*Value, error) {
	res := []*Value{}

	for _, e := range es {
		v, err := Eval(e, s)
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}

	return res, nil
}

func Eval(e *Expression, s Scope) (*Value, error) {
	if e.Val != nil {
		return e.Val, nil
	}

	if e.Var != nil {
		v, ok := s[*e.Var]
		if !ok {
			return nil, fmt.Errorf("undefined symbol: %s", *e.Var)
		}
		return v, nil
	}

	if e.Symbol != nil {
		f, ok := symbols[*e.Symbol]
		if !ok {
			return nil, fmt.Errorf("unknown symbol: %s", *e.Symbol)
		}
		return f, nil
	}

	if len(e.Call) > 0 {
		f, err := Eval(e.Call[0], s)
		if err != nil {
			return nil, err
		}

		if f.Func == nil {
			return nil, fmt.Errorf("%s is not a function", e.Call[0].String())
		}

		args, err := evalAll(e.Call[1:], s)
		if err != nil {
			return nil, err
		}

		return f.Func(args, s)
	}

	return nil, fmt.Errorf("unknown thing: %#v", e)
}
