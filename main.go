package main

import (
	"encoding/json"
	"go-neural_network/node"
)

var source = "data.json"
var memory = "memory.json"

func main() {

	brain := node.NewBrain(source, memory)
	for _, entity := range brain.GetSourceList() {
		brain.InitNextSource(node.CreateBook(entity))
		brain.Process()
	}
	q := Qwe{}
	json.Unmarshal([]byte(`{ "result": "000222a87f71972df9992f0bd9d133fb" }`), q)

}

type Qwe struct {
	result string `json:"result"`
}
