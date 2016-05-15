package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func StorePing(w http.ResponseWriter, r *http.Request) {

}

func RetrievePing(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.URL.Path)
    vars := mux.Vars(r)
    id := vars["deviceId"]
    fmt.Fprintf(w, "\ndeviceId: %v", id)
    from := vars["from"]
    fmt.Fprintf(w, "\nfrom: %v", from)
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
}

func DeleteAllPings(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Deleted all pings")
}

