package main

import (
    "neural_network/node"
)

var source = "data.json"
var memory = "memory.json"

func main() {

    brain := node.NewBrain(source, memory)
    for _, entity := range brain.GetSourceList(){
        brain.InitNextSource(node.CreateBook(),entity)
    }
    brain.Process()
}
//FreeTimeSpending, inrussianlanguage,prose, LoveLine