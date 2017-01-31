package tests

import (
	"fmt"
	"testing"
	. "github.com/hunan-rostomyan/go-index/main"
)

func dummyData() []byte {
		data := []byte(`Empiricists are in general rather suspicious with respect
to any kind of abstract entities like properties, classes, relations, numbers,
propositions, etc. They usually feel much more in sympathy with nominalists than
with realists (in the medieval sense). As far as possible they try to avoid any
reference to abstract entities and to restrict themselves to what is sometimes
called a nominalistic language, i.e., one not containing such references.`)
		return data
}

func singleLineData() []byte {
		data := []byte("Empiricists are in general rather suspicious with respect to any kind of abstract entities like properties, classes, relations, numbers, propositions, etc. They usually feel much more in sympathy with nominalists than with realists (in the medieval sense). As far as possible they try to avoid any reference to abstract entities and to restrict themselves to what is sometimes called a nominalistic language, i.e., one not containing such references.")
		return data
}

func TestTextToTokensLineNumbers(t *testing.T) {
	data := dummyData()
	tokens := TextToTokens(string(data))

	// The word 'abstract' should be on lines 1 and 4 (0-indexed).
	for _, token := range tokens {
		if token.Datum == "abstract" {
			if !(token.LineNum == 1 || token.LineNum == 4) {
				t.Error(fmt.Sprintf(
					"'abstract' should appear on lines 1 and 4, not %d",
					token.LineNum))
			}
		}
	}

	// The word 'with' should be on lines 0, 2 and 3 (0-indexed).
	for _, token := range tokens {
		if token.Datum == "with" {
			if !(token.LineNum == 0 || token.LineNum == 2 || token.LineNum == 3) {
				t.Error(fmt.Sprintf(
					"'with' should appear on lines 1 and 4, not %d",
					token.LineNum))
			}
		}
	}
}

func TestTextToTokensLineNumbersSingleLine(t *testing.T) {
	data := singleLineData()
	tokens := TextToTokens(string(data))

	// Every token should be from line 0
	for _, token := range tokens {
		if token.LineNum != 0 {
			t.Error(fmt.Sprintf(
				"token '%s' was expected to be on line 0, but was on line %d",
				token.Datum, token.LineNum))
		}
	}
}

func TestTextToTokensSize(t *testing.T) {
	data := dummyData()
	tokens := TextToTokens(string(data))

	// Ensure that the right number of tokens is generated
	if len(tokens) != 55 {
		t.Error(fmt.Sprintf("there should be %d tokens not %d", 55, len(tokens)))
	}
}

func TestTextToTokensSizeSingleLine(t *testing.T) {
	data := singleLineData()
	tokens := TextToTokens(string(data))

	// Ensure that the right number of tokens is generated
	if len(tokens) != 55 {
		t.Error(fmt.Sprintf("there should be %d tokens not %d", 55, len(tokens)))
	}
}

func TestTextToTokensLowercasing(t *testing.T) {
	text := "RepLAce"
	expected := "replace"
	actual := TextToTokens(text)[0].Datum
	if actual != expected {
		t.Error(fmt.Sprintf(
			"expected '%s' to transform to '%s' not '%s'", text, expected, actual))
	}
}
