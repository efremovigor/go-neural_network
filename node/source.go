package node

import (
	"encoding/json"
	"go-neural_network/lib"
)

type Source struct {
	path   string
	entity struct {
		Data []DataEntity `json:"data"`
	}
}

type DataEntity struct {
	Properties map[string]float64 `json:"properties"`
	Result     float64            `json:"result"`
}

func (source *Source) initDataSource(path string) {
	source.path = path
	source.setEntity(lib.ReadFile(source.path))
}

func (source *Source) setEntity(data []byte) {
	err := json.Unmarshal(data, &source.entity)
	if err != nil {
		panic(err)
	}
}
