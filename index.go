package main

import (
	"bytes"
	"fmt"
	"encoding/json"
)

type Index map[Token]map[string]float64

// From https://gist.github.com/mdwhatcott/8dd2eef0042f7f1c0cd8
func (index Index) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	size := len(index)
	i := 0
	for key, value := range index {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s", key, jsonValue))
		i += 1
		if i < size {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
