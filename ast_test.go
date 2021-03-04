package scheme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	x, err := parse("(foo bar)") // 5 3)")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(x.Call))
	assert.Equal(t, "foo", *x.Call[0].Var)
}

func Test2(t *testing.T) {
	x, err := parse("(foo 5 3.3)")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(x.Call))
	assert.Equal(t, "foo", *x.Call[0].Var)
	assert.Equal(t, 5, *x.Call[1].Val.Int)
	assert.Equal(t, 3.3, *x.Call[2].Val.Float)
}

func Test3(t *testing.T) {
	x, err := parse("(+ 5 3.3)")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(x.Call))
	assert.Equal(t, "+", *x.Call[0].Symbol)
	assert.Equal(t, 5, *x.Call[1].Val.Int)
	assert.Equal(t, 3.3, *x.Call[2].Val.Float)
}

func Test4(t *testing.T) {
	x, err := parse("(+ 5 (+1 4))")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(x.Call))
	assert.Equal(t, "+", *x.Call[0].Symbol)
	assert.Equal(t, 5, *x.Call[1].Val.Int)
	assert.Equal(t, "+", *x.Call[2].Call[0].Symbol)
}
