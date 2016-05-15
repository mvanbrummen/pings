package main

import (
    "fmt"
    "log"
    "net/http"
    //"html"

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
    to := vars["to"]
    fmt.Fprintf(w, "\nto: %v", to)
}

func RetrieveDevices(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Get all devices")
}

func DeleteAllPings(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Deleted all pings")
}

func main() {
    // initialise router
    router := mux.NewRouter().StrictSlash(true)
   
    // serve up static html
    router.Handle("/", http.FileServer(http.Dir("./static")))

    // routes
    router.HandleFunc("/test", StorePing)
    router.HandleFunc("/{deviceId}/{from}/{to}", RetrievePing)    
    router.HandleFunc("/devices", RetrieveDevices)
    router.HandleFunc("/clear_data", DeleteAllPings)

    // start server
    log.Fatal(http.ListenAndServe(":3000", router))
}

