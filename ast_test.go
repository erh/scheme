package scheme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAST0(t *testing.T) {
	x, err := Parse("1")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1.0, *x.Val.Float)

	x, err = Parse("1.1")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1.1, *x.Val.Float)

	x, err = Parse("-1.1")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, -1.1, *x.Val.Float)

	x, err = Parse("-1")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, -1.0, *x.Val.Float)

	x, err = Parse("\"a\"")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "a", *x.Val.AString)

	x, err = Parse("a")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "a", *x.Var)

	x, err = Parse("+")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "+", *x.Symbol)

	x, err = Parse("*")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "*", *x.Symbol)
}

func TestAST1(t *testing.T) {
	x, err := Parse("(foo bar)")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(x.Call))
	assert.Equal(t, "foo", *x.Call[0].Var)
}

func TestAST2(t *testing.T) {
	x, err := Parse("(foo 5 3.3)")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(x.Call))
	assert.Equal(t, "foo", *x.Call[0].Var)
	assert.Equal(t, 5.0, *x.Call[1].Val.Float)
	assert.Equal(t, 3.3, *x.Call[2].Val.Float)
}

func TestAST3(t *testing.T) {
	x, err := Parse("(+ 5 3.3)")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(x.Call))
	assert.Equal(t, "+", *x.Call[0].Symbol)
	assert.Equal(t, 5.0, *x.Call[1].Val.Float)
	assert.Equal(t, 3.3, *x.Call[2].Val.Float)
}

func TestAST4(t *testing.T) {
	x, err := Parse("(+ 5 (+ 1 4))")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(x.Call))
	assert.Equal(t, "+", *x.Call[0].Symbol)
	assert.Equal(t, 5.0, *x.Call[1].Val.Float)
	assert.Equal(t, "+", *x.Call[2].Call[0].Symbol)
}
