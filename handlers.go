package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	DataFile   string = "ping.data"
	AllDevices string = "all"
)

func StorePing(w http.ResponseWriter, r *http.Request) {
	// get path params
	vars := mux.Vars(r)
	id := vars["deviceId"]
	time := vars["epochTime"]
	if IsDateFormat(FmtUnix, time) {
		t, _ := strconv.ParseInt(time, 0, 64)

		// get data from file otherwise instantiate
		var m map[string]int64arr
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
			log.Print("Failed to write data.")
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Epoch time was not a unix timestamp", http.StatusBadRequest)
	}
}

func RetrievePing(w http.ResponseWriter, r *http.Request) {
	// get path params
	vars := mux.Vars(r)
	id := vars["deviceId"]
	from := vars["from"]
	if IsDateFormat(FmtISO, from) {
		m, err := GetPingMap()
		if err != nil {
			log.Print("Could not get ping map.")
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}

		if id == AllDevices {
			// filter on every slice
			devices, err := GetPingMapKeys()
			if err != nil {
				log.Print("Failed to get ping map keys.")
				http.Error(w, "Internal server error.", http.StatusInternalServerError)
			}
			for _, k := range devices {
				m[k] = m[k].Pings(from)
			}
			pingResponse, err := json.Marshal(m)
			if err != nil {
				log.Print("Failed to marshal to json.")
				http.Error(w, "Internal server error.", http.StatusInternalServerError)
			}
			fmt.Fprintf(w, "%s", pingResponse)
		} else {
			// filter on the specified slice
			m[id] = m[id].Pings(from)
			pingResponse, err := json.Marshal(m[id])
			if err != nil {
				log.Print("Failed to marshal to json.")
				http.Error(w, "Internal server error.", http.StatusInternalServerError)
			}
			fmt.Fprintf(w, "%s", pingResponse)
		}
	} else {
		http.Error(w, "From was not an ISO date.", http.StatusBadRequest)
	}
}

func RetrievePingTo(w http.ResponseWriter, r *http.Request) {
	// get path params
	vars := mux.Vars(r)
	id := vars["deviceId"]
	from := vars["from"]
	to := vars["to"]

	m, err := GetPingMap()
	if err != nil {
		log.Print("Could not get ping map.")
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}
	if id == AllDevices {
		// filter on every slice
		devices, err := GetPingMapKeys()
		if err != nil {
			log.Print("Failed to get ping map keys.")
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
		for _, k := range devices {
			m[k] = m[k].PingsTo(from, to)
		}
		pingResponse, err := json.Marshal(m)
		if err != nil {
			log.Print("Failed to marshal to json.")
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "%s", pingResponse)
	} else {
		// filter on the specified slice
		m[id] = m[id].PingsTo(from, to)
		pingResponse, err := json.Marshal(m[id])
		if err != nil {
			log.Print("Failed to marshal to json.")
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "%s", pingResponse)
	}
}

func RetrieveDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := GetPingMapKeys()
	if err != nil {
		log.Print("Failed to get ping map keys.")
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}
	devicesResponse, err := json.Marshal(devices)
	if err != nil {
		log.Print("Failed to marshal to json.")
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%s", devicesResponse)
}

func DeleteAllPings(w http.ResponseWriter, r *http.Request) {
	if PingMapCreated() {
		err := os.Remove(DataFile)
		if err != nil {
			log.Print("Failed to delete ping data file.")
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
		}
	}
}
