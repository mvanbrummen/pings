package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strconv"
    "os"

    "github.com/gorilla/mux"
)

const (
    DataFile string = "ping.data"
)

func StorePing(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["deviceId"]
    time := vars["epochTime"]
    t, _ := strconv.ParseUint(time, 0, 32)

    var m map[string][]uint32;
    if PingMapCreated() {
   	m, _ = GetPingMap()
    } else {
        m = make(map[string][]uint32)
    }
    m[id] = append(m[id], uint32(t))
    b, _ := GetBytes(m)
    err := ioutil.WriteFile(DataFile, b, 0600)  
    if err != nil {
        fmt.Printf("Failed to write data.")
    }
}

func RetrievePing(w http.ResponseWriter, r *http.Request) {
    //vars := mux.Vars(r)
    //id := vars["deviceId"]
    //from := vars["from"]

    m, err := GetPingMap()
    if err != nil {
	fmt.Printf("Could not get ping map")
    }
    //pings := m[id]
    pingResponse, err := json.Marshal(m)
    fmt.Fprintf(w, "%s", pingResponse)
}

func RetrievePingTo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.URL.Path)
    vars := mux.Vars(r)
    id := vars["deviceId"]
    fmt.Fprintf(w, "\ndeviceId: %v", id)
    from := vars["from"]
    fmt.Fprintf(w, "\nfrom: %v", from)
    to := vars["to"]
    fmt.Fprintf(w, "\nto: %v", to)
}

func RetrieveDevices(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Get all devices")
    devices, err := GetPingMapKeys()
    if err != nil {
	fmt.Printf("Failed to get ping map keys.")
    }
    devicesResponse, _ := json.Marshal(devices)
    fmt.Fprintf(w, "%s", devicesResponse) 
}

func DeleteAllPings(w http.ResponseWriter, r *http.Request) {
    err := os.Remove(DataFile)     
    if err != nil {
 	fmt.Printf("Failed to delete ping data file.")
    }   
}

