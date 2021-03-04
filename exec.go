package scheme

import (
	"fmt"
)

type Scope map[string]*Value

type SchemeFunction func(args []*Value, scope Scope) (*Value, error)

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
			v, ok = builtins[*e.Var]
		}

		if !ok {
			return nil, fmt.Errorf("undefined variable: %s", *e.Var)
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
