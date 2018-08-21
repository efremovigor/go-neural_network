package node

type FormInterface interface {
    Init(list []string)
    getProperties() map[string]*Neuron
    setDefault(list []string,key string)
}

type Form struct {
    Properties map[string]*Neuron
    Result     bool
}

func (form *Form) Init(list []string) {
    form.Properties = make(map[string]*Neuron)
    for _, item := range list {
        form.Properties[item] = createNeuron(item)
    }
}

func (form *Form) getProperties() map[string]*Neuron {
    return form.Properties
}

func (form *Form) setDefault(list []string,key string)  {
    set := false
    for _,value := range list {
        if set == true {
            form.Properties[value].State = false
            continue
        }
        if form.Properties[value].State == true {
            set = true
            continue
        }
    }
    if set == false {
        form.Properties[key].State = true
    }
}
