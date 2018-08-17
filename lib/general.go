package lib

import "io/ioutil"

func ReadFile(path string) []byte {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
    return data
}
