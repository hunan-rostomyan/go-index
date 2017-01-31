package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	. "github.com/hunan-rostomyan/go-index/main"
)

func createDummyData() string {
	dir, _ := ioutil.TempDir(".", "tmp")
	data1 := []byte("first\nsecond\nthird\nfourth line\n")
	data2 := []byte("how\nare\nyou\ndoing?\n")
	data3 := []byte("BSD\n")
	_ = ioutil.WriteFile(filepath.Join(dir, "file1.html.md.erb"), data1, 0644)
	_ = ioutil.WriteFile(filepath.Join(dir, "file2.html.md.erb"), data2, 0644)
	_ = ioutil.WriteFile(filepath.Join(dir, "LICENSE"), data3, 0644)
	return dir
}

func createDummyDocument() (string, []byte) {
	dir, _ := ioutil.TempDir(".", "tmp")
	data := []byte(`Empiricists are in general rather suspicious with respect
to any kind of abstract entities like properties, classes, relations, numbers,
propositions, etc. They usually feel much more in sympathy with nominalists than
with realists (in the medieval sense). As far as possible they try to avoid any
reference to abstract entities and to restrict themselves to what is sometimes
called a nominalistic language, i.e., one not containing such references.`)
	_ = ioutil.WriteFile(filepath.Join(dir, "carnap.excerpt"), data, 0644)
	return dir, data
}

func TestPatternAll(t *testing.T) {
	dir := createDummyData()
	defer os.RemoveAll(dir)

	col := NewLocalCollection(dir, "(.*)")
	_, docNum := col.Documents()
	if docNum != 3 {
		t.Error(fmt.Sprintf("expected 3 documents found %d", docNum))
	}
}

func TestPatternEmptyString(t *testing.T) {
	dir := createDummyData()
	defer os.RemoveAll(dir)

	// All file names should match the empty string.
	col := NewLocalCollection(dir, "")
	_, docNum := col.Documents()
	if docNum != 3 {
		t.Error(fmt.Sprintf("expected 3 documents found %d", docNum))
	}
}

func TestPatternEndsWith(t *testing.T) {
	dir := createDummyData()
	defer os.RemoveAll(dir)

	col := NewLocalCollection(dir, ".html.md.erb$")
	_, docNum := col.Documents()
	if docNum != 2 {
		t.Error(fmt.Sprintf("expected 2 documents found %d", docNum))
	}
}

func TestPatternExactMatch(t *testing.T) {
	dir := createDummyData()
	defer os.RemoveAll(dir)

	col := NewLocalCollection(dir, "^LICENSE$")
	_, docNum := col.Documents()
	if docNum != 1 {
		t.Error(fmt.Sprintf("expected 1 document found %d", docNum))
	}
}

func TestDocumentContents(t *testing.T) {
	dir, data := createDummyDocument()
	defer os.RemoveAll(dir)

	col := NewLocalCollection(dir, "^carnap.excerpt$")
	docs, _ := col.Documents()
	doc := docs[0]

	// Ensure document contents is the right list of tokens.
	expectedTokens := TextToTokens(string(data))
	expectedTokenNum := len(expectedTokens)
	actualTokens := doc.Contents
	actualTokenNum := len(actualTokens)
	if actualTokenNum != expectedTokenNum {
		t.Error(fmt.Sprintf("expected size of tokens to be %d was %d",
			expectedTokenNum, actualTokenNum))
	}
}
