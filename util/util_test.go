package util

import (
	"fmt"
	"testing"
)

type TStringer struct {
	val string
	fmt.Stringer
}

func (ts *TStringer) String() string {
	return ts.val
}

func TestConcat(t *testing.T) {
	stringer1 := &TStringer{val: "foo"}
	stringer2 := &TStringer{val: "bar"}

	if expected, actual := "foo", Concat(stringer1); actual != expected {
		t.Errorf("expected %s; got %s", expected, actual)
	}

	if expected, actual := "foobar", Concat(stringer1, stringer2); actual != expected {
		t.Errorf("expected %s; got %s", expected, actual)
	}
}
