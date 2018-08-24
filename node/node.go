package node

import (
    "neural_network/lib"
    "encoding/json"
    "crypto/md5"
    "encoding/hex"
    "time"
    "strconv"
)

type Neuron struct {
    Id    string
    Name  string
    State bool
}

func createNeuron(name string) *Neuron {
    h := md5.New()
    h.Write([]byte(name + strconv.FormatInt(time.Now().Unix(),10)))
    id := hex.EncodeToString(h.Sum(nil))

    return &Neuron{State: false, Name: name, Id:id[:4]}
}

type Brain struct {
    pathSource     string
    pathMemory     string
    memory         map[string]float64
    CurrentProcess BrainProcess
    Source         Source
}

type BrainProcess struct {
    Form         FormInterface
    Neuron       map[string]Neuron
    hideNeuron   map[string]Neuron
    weight       map[string]float64
    resultNeuron Neuron
}

func (b *Brain) init() {
    b.Source = Source{}
    b.Source.initDataSource(b.pathSource)
    b.initMemory()
}

func (b *Brain) newProcess(form FormInterface) {
    b.CurrentProcess = BrainProcess{Form: form}
}

func (b *Brain) initMemory() {
    json.Unmarshal(lib.ReadFile(b.pathMemory), &b.memory)
}

func (b *Brain) GetSourceList() []DataEntity {
    return b.Source.entity.Data
}

func (b *Brain) Process() {
    //b.CurrentProcess.Form.AutoSetProperties()
}

func NewBrain(pathSource string, pathMemory string) (brain Brain) {
    brain = Brain{pathSource: pathSource, pathMemory: pathMemory}
    brain.init()
    return
}

func (b *Brain) InitNextSource(form FormInterface, entity DataEntity) {
    b.newProcess(form)
    for key, value := range entity.Properties {
        form.SetProperty(key, value)
    }
}
