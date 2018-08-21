package node

import (
    "encoding/json"
    "neural_network/lib"
)

type Source struct {
    path   string
    entity SourceEntity
}

type SourceEntity struct {
    data []DataEntity `json:"data"`
}

type DataEntity struct {
    Properties map[string]bool `json:"properties"`
    Result     bool            `json:"result"`
}


func (source Source) initDataSource(path string) {
    source.path = path
    source.setEntity(lib.ReadFile(source.path))
}

func (source Source) setEntity(data []byte) {
    err := json.Unmarshal(data,&source.entity)
    if err != nil {
        panic(err)
    }
}
