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
	index := BuildIndex(col.Documents())

	// Save
	b, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	_ = ioutil.WriteFile("index.json", b, 0644)
}
