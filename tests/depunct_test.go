package tests

import (
	"fmt"
	"testing"
	. "github.com/hunan-rostomyan/go-index/main"
)

func TestDepunctLowercaseAlphaOk(t *testing.T) {
	text := "rudolf carnap"
	expecting := "rudolf carnap"
	actual := Depunct(text)
	if actual !=  expecting{
		t.Error(fmt.Sprintf("was expecting '%s', got '%s'", expecting, actual))
	}
}

func TestDepunctDashesOk(t *testing.T) {
	text := "self-signed"
	expecting := "self-signed"
	actual := Depunct(text)
	if actual !=  expecting{
		t.Error(fmt.Sprintf("was expecting '%s', got '%s'", expecting, actual))
	}
}

func TestDepunctAllPunct(t *testing.T) {
	text := "ALL.?=+"
	expecting := "A      "  // 'A' is a hex digit
	actual := Depunct(text)
	if actual !=  expecting{
		t.Error(fmt.Sprintf("was expecting '%s', got '%s'", expecting, actual))
	}
}
