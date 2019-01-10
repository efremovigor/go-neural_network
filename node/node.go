package node

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"go-neural_network/lib"
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
	Value int64
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
  CurrentProcess - текущий процесс прогона сети
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

	CurrentProcess BrainProcess
	Source         Source
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

func (process *BrainProcess) initMemoryNeuron() {

}

func (process *BrainProcess) initInputNeural() {
	for _, value := range process.Form.GetProperties() {
		if val, ok := process.input[value.Id]; ok {
			//do something here
		}

	}
}

func (this *Brain) init() {
	this.Source = Source{}
	this.Source.initDataSource(this.pathSource)
	this.initMemory()
}

func (this *Brain) newProcess(form FormInterface) {
	this.CurrentProcess = BrainProcess{Form: form}

}

func (this *Brain) initMemory() {
	json.Unmarshal(lib.ReadFile(this.pathMemory), &this.memory)
}

func (this *Brain) GetSourceList() []DataEntity {
	return this.Source.entity.Data
}

func (this *Brain) Process() {
	this.CurrentProcess.Form.AutoSetProperties()
	this.CurrentProcess.initInputNeural()
}

func NewBrain(pathSource string, pathMemory string) (brain Brain) {
	brain = Brain{pathSource: pathSource, pathMemory: pathMemory}
	brain.init()
	return
}

func (this *Brain) InitNextSource(form FormInterface) {
	this.newProcess(form)
	this.CurrentProcess.initMemoryNeuron()

}
