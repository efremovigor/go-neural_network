package node

import "go-neural_network/girlForm"

type Girl struct {
	Form
}

func CreateGirl() (girl Girl) {
	girl.Init(girlForm.CollectionProperties)
	return
}

func (girl *Girl) SetHair(key string) {
	for _, value := range girlForm.TitsProperties {
		if value == key {

			continue
		}
		if neuron, ok := girl.Properties[value]; ok {
			neuron.Value = 1
		}
	}
}
