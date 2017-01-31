package tests

import (
	"testing"
	. "github.com/hunan-rostomyan/go-index/main"
)

func TestStop(t *testing.T) {
	tokens := []Token{
		Token{"the", 0},
		Token{"of", 0},
	}
	if len(Stop(tokens)) != 0 {
		t.Error("both tokens are stops, should the list should be empty")
	}
}
