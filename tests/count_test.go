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
	counts := Count(tokens)
	for token := range counts {
		switch token.Datum {
		case "foo":
			if counts[token] != 1 {
				t.Error(fmt.Sprintf("expecting 'foo' count to be 1 not %d", counts[token]))
			}
		case "bar":
			if counts[token] != 2 {
				t.Error(fmt.Sprintf("expecting 'bar' count to be 2 not %d", counts[token]))
			}
		case "baz":
			if counts[token] != 3 {
				t.Error(fmt.Sprintf("expecting 'baz' count to be 3 not %d", counts[token]))
			}
		}
	}
}
