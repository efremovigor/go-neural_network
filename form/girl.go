package form

type Girl struct {
    Form
}

func CreateGirl() (girl Girl) {
    girl.init(girlCollectionProperties)
    return
}


func (girl *Girl) SetHair(key string) {
    for _, value := range titsProperties {
        if value == key {

            continue
        }
        if neuron, ok := girl.properties[value]; ok {
            neuron.state = true
        }
    }
}
