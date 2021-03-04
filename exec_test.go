package scheme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func evalTest(t *testing.T, s string, correct interface{}, scope Scope) {
	x, err := parse(s)
	if err != nil {
		t.Fatal(err)
	}

	y, err := Eval(x, scope)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, correct, y.Primitive())
}

func TestPriimitives(t *testing.T) {
	scope := Scope{}

	evalTest(t, "5", 5, scope)
	evalTest(t, "5.3", 5.3, scope)
	evalTest(t, "\"abc\"", "abc", scope)
}

func TestMath1(t *testing.T) {
	scope := Scope{}
	evalTest(t, "(+ 1 2.1)", 3.1, scope)
	evalTest(t, "(+ 1 2.1 1.1)", 4.2, scope)

	evalTest(t, "(* 2 2.1)", 4.2, scope)
	evalTest(t, "(* 2 2.1 2)", 8.4, scope)

	v := 1501.0
	scope["raw"] = &Value{Float: &v}
	evalTest(t, "(- raw 1100)", 401.0, scope)

	evalTest(t, "(/ (- raw 1100) 10)", 40.1, scope)
}
