package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strconv"
    "os"
    "sort"
    "time"

    "github.com/gorilla/mux"
)

const (
    DataFile string = "ping.data"
    AllDevices string = "all"
    DateLayoutISO string = "2006-01-02"
)

func StorePing(w http.ResponseWriter, r *http.Request) {
    // get path params
    vars := mux.Vars(r)
    id := vars["deviceId"]
    time := vars["epochTime"]
    t, _ := strconv.ParseInt(time, 0, 64)

    // get data from file otherwise instantiate
    var m map[string]int64arr;
    if PingMapCreated() {
   	m, _ = GetPingMap()
    } else {
        m = make(map[string]int64arr)
    }

    // append time to slice
    m[id] = append(m[id], int64(t))

    // sort slice
    sort.Sort(m[id])

    // write data to file
    b, _ := GetBytes(m)
    err := ioutil.WriteFile(DataFile, b, 0600)  
    if err != nil {
        fmt.Printf("Failed to write data.")
    }
}

func RetrievePing(w http.ResponseWriter, r *http.Request) {
    // get path params
    vars := mux.Vars(r)
    id := vars["deviceId"]
    from := vars["from"]

    // parse from to ISO date
    fromDate, err := time.Parse(DateLayoutISO, from) 
    if err != nil {
 	fmt.Printf("Date parse error")
    }
    fmt.Printf("ISO date is %v", fromDate)

    m, err := GetPingMap()
    if err != nil {
	fmt.Printf("Could not get ping map")
    }

    if id == AllDevices {
	// filter on every slice	
    } else {
        pings := m[id]
        pingResponse, err := json.Marshal(pings)
    	if err != nil {
	    fmt.Printf("Failed to marshal to json")
        }
        fmt.Fprintf(w, "%s", pingResponse)
    }
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
    devices, err := GetPingMapKeys()
    if err != nil {
	fmt.Printf("Failed to get ping map keys.")
    }
    devicesResponse, _ := json.Marshal(devices)
    fmt.Fprintf(w, "%s", devicesResponse) 
}

func DeleteAllPings(w http.ResponseWriter, r *http.Request) {
    if PingMapCreated() {
        err := os.Remove(DataFile)     
        if err != nil {
 	    fmt.Printf("Failed to delete ping data file.")
        }   
    }
}

