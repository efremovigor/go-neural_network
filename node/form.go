package node

type FormInterface interface {
	Init(list []string)
	GetProperties() map[string]*Neuron
	setDefault(list []string, key string)
	SetProperty(name string, value float64)
	SetResult(value float64)
	GetResult() float64
	AutoSetProperties()
}

type Form struct {
	Properties map[string]*Neuron
	Result     float64
}

func (form *Form) Init(list []string) {
	form.Properties = make(map[string]*Neuron)
	for _, item := range list {
		form.Properties[item] = createNeuron(item)
	}
}

func (form *Form) GetProperties() map[string]*Neuron {
	return form.Properties
}

func (form *Form) setDefault(list []string, key string) {
	set := false
	for _, value := range list {
		if set == true {
			form.Properties[value].Value = 0
			continue
		}
		if form.Properties[value].Value == 1 {
			set = true
			continue
		}
	}
	if set == false {
		form.Properties[key].Value = 1
	}
}

func (form *Form) SetProperty(name string, value float64) {
	form.Properties[name].Value = value
}

func (form *Form) SetResult(value float64) {
	form.Result = value
}

func (form *Form) GetResult() float64 {
	return form.Result
}
