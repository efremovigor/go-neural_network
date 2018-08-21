package main

import (
    "neural_network/node"
)

var source = "data.json"
var memory = "memory.json"

func main() {

    brain := node.NewBrain(source, memory)
    var book node.Book
    for _, entity := range brain.GetSourceList(){
        process := node.BrainProcess{Form:node.CreateBook()}
        for key,property := range entity.Properties{
            book.Properties[key].State = property
        }

    }
    brain.Process()
}
//FreeTimeSpending, inrussianlanguage,prose, LoveLine