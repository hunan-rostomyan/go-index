package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// Collection represents a group of Documents. These documents may be physically
// located in all sorts of places, so implementations of this interface will have
// to decide how exactly these Documents are to be found.
type Collection interface {
	Documents() (docs []Document, size int)
}

// Filterable is intended to be embedded in those Collections for which it makes
// sense to filter the underlying files by matching a regular expression pattern
// against their names.
type Filterable struct {
	Pattern *regexp.Regexp
}

// LocalCollection is a Collection of Documents located on a local machine. It is
// filterable, so the pattern is used to filter in only files having specific name
// forms.
type LocalCollection struct {
	Path string
	Filterable
}

// Documents reads the contents of the collection path, filtering in those files
// that have the pattern of interest. It then reads, cleans, tokenizes, and turns
// them into Documents. It returns the list of documents and its size.
//
// TODO: cleaning and tokenization should be extracted out
// DISCUSS: consider making the traversal recursive
func (col LocalCollection) Documents() ([]Document, int) {
	NewDocument := DocumentFactory()
	docs := []Document{}
	size := 0

	files, _ := ioutil.ReadDir(col.Path)
	for _, f := range files {
		if f.Mode().IsRegular() {
			fileName := f.Name()
			if col.Pattern.MatchString(fileName) {
				contents, _ := ioutil.ReadFile(filepath.Join(col.Path, fileName))
				tokens := TextToTokens(string(contents))
				docs = append(docs, NewDocument(fileName, tokens))
				size++
			}
		}
	}
	return docs, size
}

// NewLocalCollection takes a local directory path and a regular expression
// pattern, returning a LocalCollection.
func NewLocalCollection(path string, pattern string) Collection {
	r, _ := regexp.Compile(pattern)
	return LocalCollection{
		Path:       path,
		Filterable: Filterable{r}}
}

func (col LocalCollection) String() string {
	return fmt.Sprintf("Collection ('%s')", col.Path)
}
