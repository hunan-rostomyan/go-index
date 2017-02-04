package tests

import (
	"fmt"
	"testing"
	. "github.com/hunan-rostomyan/go-index/main"
)

func TestCount(t *testing.T) {
	tokens := []Token{
		Token{"foo", 0},
		Token{"bar", 0},
		Token{"bar", 1},
		Token{"baz", 0},
		Token{"baz", 1},
		Token{"baz", 2},
	}

	counter := NewCounter(tokens)

	if counter.Lookup("foo") != 1 {
		t.Error(fmt.Sprintf("expecting 'foo' count to be 1 not %d", counter.Lookup("foo")))
	}
	if counter.Lookup("bar") != 2 {
		t.Error(fmt.Sprintf("expecting 'bar' count to be 2 not %d", counter.Lookup("bar")))
	}
	if counter.Lookup("baz") != 3 {
		t.Error(fmt.Sprintf("expecting 'baz' count to be 3 not %d", counter.Lookup("baz")))
	}

}
