package main

import (
    "fmt"
    "bytes"
    "encoding/gob"
    "io/ioutil"
    "os"
)

func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func FromBytes(data []byte, obj interface{}) (error) {
    b := bytes.NewBuffer(data)
    dec := gob.NewDecoder(b)
    err := dec.Decode(obj)
    if err != nil {
        return err
    }
    return nil
}

func GetPingMap() (map[string][]uint32, error){
    b, _ := ioutil.ReadFile(DataFile)
    m := make(map[string][]uint32)
    err := FromBytes(b, &m)
    if err != nil {
        fmt.Printf("Failed to get from bytes.")
  	return nil, err
    }
    return m, nil
}

func GetPingMapKeys() ([]string, error) {
    m, err := GetPingMap()
    if err != nil {
	return nil, err
    }
    keys := make([]string, len(m))
    i := 0
    for k := range m {
	keys[i] = k 
        i++
    }
    return keys, nil
}

func PingMapCreated() bool {
    if _, err := os.Stat(DataFile); err == nil {
        return true
    } else {
        return false
    }
}
