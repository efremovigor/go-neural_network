package node

type Neuron struct {
    id    int
    Name  string
    State bool
}

func CreateNeuron(name string) *Neuron {
    return &Neuron{State: false, Name: name}
}

type Brain struct {
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

func (b *Brain) init(path string)  {
    b.Source = Source{}
    b.Source.InitDataSource(path)
    b.newProcess()
}

func (b *Brain) newProcess()  {
    b.CurrentProcess = BrainProcess{}
}

func NewBrain(path string) (brain Brain) {
    brain = Brain{}
    brain.init(path)
    return
}

func (b *Brain) Process()  {

}
