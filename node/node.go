package node

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"go-neural_network/lib"
	"math"
	"strconv"
	"time"
)

/**
  Id - уникальный id неирона
  Name - название неирона
  Value - значение неирона
*/
type Neuron struct {
	Id    string
	Name  string
	Value float64
}

/**
  Создание пустого неирона по имени
*/
func createNeuron(name string) *Neuron {
	h := md5.New()
	h.Write([]byte(name + strconv.FormatInt(time.Now().Unix(), 10)))
	id := hex.EncodeToString(h.Sum(nil))

	return &Neuron{Value: 0, Name: name, Id: id}
}

/**
  pathSource - откуда берутся данные по обучению
  pathMemory - место хранения состояния сети
  memory - данные из памяти  - ключ неирона -> текущий вес
  Process - текущий процесс прогона сети
  Source - данные для прогона сети
*/
type Brain struct {
	pathSource string
	pathMemory string
	memory     struct {
		Weight map[string]float64 `json:"weight"`
		Neuron struct {
			Input  map[string]string `json:"input"`
			Hide   []string          `json:"hide"`
			Result string            `json:"result"`
		} `json:"neuron"`
	}

	Process *BrainProcess
	Source  Source
}

/**
  Form - все входящие данные для прогона сети
  input - неироны входящие
  hide - неироны скрытого уровня
  weight - ключи input.neuron.id + hide.neuron.id или hide.neuron.id + result.neuron.id
*/
type BrainProcess struct {
	Form   FormInterface
	input  map[string]Neuron
	hide   map[string]Neuron
	weight map[string]float64
	result Neuron
}

func (brain *Brain) initMemoryInProcess() {

}

func (process *BrainProcess) generateHideLayer(count int) {
	for i := 0; i < count; i++ {
		newNeuron := createNeuron(lib.RandStringBytes(8))
		process.hide[newNeuron.Id] = *newNeuron
	}
}

func (process *BrainProcess) generateWeight() {
	for inputKey := range process.input {
		for hideKey := range process.hide {
			process.weight[inputKey+"-"+hideKey] = lib.RandFloat(0.0001, 0.1)
		}
	}
	for hideKey := range process.hide {

	}
}

func (process *BrainProcess) initInputNeural() {
	for _, value := range process.Form.GetProperties() {
		if val, ok := process.input[value.Name]; ok {
			if value.Id != val.Id {
				panic("Neuron's hash are not identical")
			}
			val.Value = value.Value
		} else {
			process.input[value.Name] = *value
		}
	}
}

func (process *BrainProcess) estimationHideLayer() {
	var weight float64
	var total float64
	for hideKey, hideNeuron := range process.hide {
		weight = 0
		total = 0
		for inputKey, inputNeuron := range process.input {
			if _, ok := process.weight[inputKey+"-"+hideKey]; ok {
				weight = process.weight[inputKey+"-"+hideKey]
			} else {
				panic("Invalid the key of weight")
			}
			total += weight * float64(inputNeuron.Value)
		}
		hideNeuron.Value = math.Floor(1/(1+math.Pow(math.E, -total))*10000) / 10000
		process.hide[hideKey] = hideNeuron
	}
}

func (brain *Brain) init() {
	brain.Source = Source{}
	brain.Source.initDataSource(brain.pathSource)
	brain.initMemory()
}

func (brain *Brain) getCurrentProcess() *BrainProcess {
	return brain.Process
}

func (brain *Brain) newProcess(form FormInterface) {
	brain.Process = &BrainProcess{Form: form, input: map[string]Neuron{}, hide: map[string]Neuron{}, weight: map[string]float64{}}

}

func (brain *Brain) initMemory() {
	_ = json.Unmarshal(lib.ReadFile(brain.pathMemory), &brain.memory)
}

func (brain *Brain) GetSourceList() []DataEntity {
	return brain.Source.entity.Data
}

func (brain *Brain) Execute() {
	brain.getCurrentProcess().estimationHideLayer()
}

func NewBrain(pathSource string, pathMemory string) (brain Brain) {
	brain = Brain{pathSource: pathSource, pathMemory: pathMemory}
	brain.init()
	return
}

func (brain *Brain) InitNextSource(form FormInterface) {
	brain.newProcess(form)
	brain.initMemoryInProcess()
	brain.getCurrentProcess().Form.AutoSetProperties()
	brain.getCurrentProcess().initInputNeural()

	if len(brain.getCurrentProcess().hide) == 0 {
		brain.getCurrentProcess().generateHideLayer(len(brain.getCurrentProcess().input))
	}

	if brain.getCurrentProcess().result.Id == "" {
		brain.getCurrentProcess().result = *createNeuron(lib.RandStringBytes(8))
	}

	if len(brain.getCurrentProcess().weight) == 0 {
		brain.getCurrentProcess().generateWeight()
	}


}
