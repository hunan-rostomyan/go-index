package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("../../cloudfoundry/docs-bosh")
	col := NewLocalCollection(path, "(.*).html.md.erb$")

	// Build the inverted index
	invertedIndex := BuildInvertedIndex(col.Documents())
	b, err := json.MarshalIndent(invertedIndex, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	_ = ioutil.WriteFile("index.json", b, 0644)
}
