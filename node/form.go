package node

type Form struct {
    Properties map[string]*Neuron
    Result     bool
}

func (form *Form) Init(list []string) {
    form.Properties = make(map[string]*Neuron)
    for _, item := range list {
        form.Properties[item] = CreateNeuron(item)
    }
}

func (form *Form) GetProperties() map[string]*Neuron {
    return form.Properties
}
