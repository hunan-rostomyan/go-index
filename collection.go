package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type Collection interface {
	Documents() (docs []Document)
}

type FilterableCollection struct {
	Pattern *regexp.Regexp
}

type LocalCollection struct {
	Location string
	FilterableCollection
}

func (col LocalCollection) Documents() []Document {
	NewDocument := DocumentFactory()
	docs := []Document{}

	files, _ := ioutil.ReadDir(col.Location)
	for _, f := range files {
		if f.Mode().IsRegular() {
			fileName := f.Name()
			if col.Pattern.MatchString(fileName) {
				contents, _ := ioutil.ReadFile(filepath.Join(col.Location, fileName))
				// Lowercase
				contentsStr := strings.ToLower(string(contents))
				// Remove punctuation
				contentsStr = Depunct(contentsStr)
				// Tokenize
				tokens := Tokenize(contentsStr)
				// Stop
				tokens = Stop(tokens)
				docs = append(docs, NewDocument(fileName, tokens))
			}
		}
	}
	return docs
}

func NewLocalCollection(location string, pattern string) Collection {
	r, _ := regexp.Compile(pattern)
	return LocalCollection{
		Location: location,
		FilterableCollection: FilterableCollection{r}}
}

func (col LocalCollection) String() string {
	return fmt.Sprintf("Collection ('%s')", col.Location)
}
