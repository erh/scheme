package scheme

import (
	"fmt"
)

var (
	symbols  = map[string]*Value{}
	builtins = map[string]*Value{}
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

	builtins["max"] = &Value{Func: func(args []*Value, scope Scope) (*Value, error) {
		if len(args) == 0 {
			return nil, fmt.Errorf("max needs at least 1 argument")
		}

		max := 0.0

		for idx, a := range args {
			n, err := a.ToFloat()
			if err != nil {
				return nil, fmt.Errorf("arg %d to max fail: %s", idx, err)
			}

			if idx == 0 {
				max = n
			} else if n > max {
				max = n
			}
		}

		return &Value{Float: &max}, nil

	}}
}
