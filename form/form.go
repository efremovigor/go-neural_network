package form

type Neuron struct {
    id    int
    name  string
    state bool
}

type HideNeuron struct {
    weight    float64
    weightIn  map[int]*Neuron
    weightOut float64
}

func createNeuron(name string) *Neuron {
    return &Neuron{state: false, name: name}
}

type Form struct {
    properties map[string]*Neuron
    result     bool
}

func (form *Form) init(list []string) {
    form.properties = make(map[string]*Neuron)
    for _, item := range list {
        form.properties[item] = createNeuron(item)
    }
}

func (form *Form) GetProperties() map[string]*Neuron {
    return form.properties
}
