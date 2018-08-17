package main

import (
    "neural_network/node"
)

func main() {

    brain := node.NewBrain("data.json","memory.json")
    brain.Process()
}