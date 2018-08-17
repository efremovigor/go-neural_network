package node

import (
    "neural_network/lib"
    "encoding/json"
)

type Neuron struct {
    id    int
    Name  string
    State bool
}

func createNeuron(name string) *Neuron {
    return &Neuron{State: false, Name: name}
}

type Brain struct {
    pathSource     string
    pathMemory     string
    memory         map[string]float64
    CurrentProcess BrainProcess
    Source         Source
}

type BrainProcess struct {
    entity       Form
    Neuron       map[string]Neuron
    hideNeuron   map[string]Neuron
    weight       map[string]float64
    resultNeuron Neuron
}

func (b *Brain) init() {
    b.Source = Source{}
    b.Source.initDataSource(b.pathSource)
    b.initMemory()
    b.newProcess()
}

func (b *Brain) newProcess() {
    b.CurrentProcess = BrainProcess{}
}

func (b *Brain) initMemory() {
    json.Unmarshal(lib.ReadFile(b.pathMemory),&b.memory)
}

func NewBrain(pathSource string, pathMemory string) (brain Brain) {
    brain = Brain{pathSource: pathSource, pathMemory: pathMemory}
    brain.init()
    return
}

func (b *Brain) Process() {

}
