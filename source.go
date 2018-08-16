package main

import (
    "io/ioutil"
    "encoding/json"
)

type Source struct {
    path   string
    entity SourceEntity
}

type SourceEntity struct {
    Data DataListEntity `json:"data"`
}

type DataListEntity []DataEntity


type DataEntity struct {
    Properties map[string]bool `json:"properties"`
    Result     bool            `json:"result"`
}


func (source Source) initDataSource() {
    source.path = sourcePath
    source.setEntity(source.loadFile())
}

func (source Source) loadFile() []byte {
    data, err := ioutil.ReadFile(source.path)
    if err != nil {
        panic(err)
    }
    return data
}

func (source Source) setEntity(data []byte) {
    err := json.Unmarshal(data,&source.entity)
    if err != nil {
        println(err)
    }
}
