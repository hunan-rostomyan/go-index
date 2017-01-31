package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	col := NewLocalCollection("/Users/hunan/experiments/index/docs-bosh", "(.*).html.md.erb$")
	index := BuildIndex(col.Documents())

	// Save
	b, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	_ = ioutil.WriteFile("index.json", b, 0644)
}
