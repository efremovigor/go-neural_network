package main

import (
    "neural_network/node"
)

var source = "data.json"
var memory = "memory.json"

func main() {

    brain := node.NewBrain(source, memory)
    brain.Process()
}
