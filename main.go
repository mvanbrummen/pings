package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

const (
    Post string = "POST"
    Get string = "GET"
)

func main() {
    // initialise router
    router := mux.NewRouter().StrictSlash(true)

    // routes
    router.HandleFunc("/{deviceId}/{epochTime}", StorePing).Methods(Post)
    router.HandleFunc("/{deviceId}/{from}", RetrievePing).Methods(Get)    

    router.HandleFunc("/{deviceId}/{from}/{to}", RetrievePingTo).Methods(Get)    
    router.HandleFunc("/devices", RetrieveDevices).Methods(Get)
    router.HandleFunc("/clear_data", DeleteAllPings).Methods(Post)

    // start server
    log.Print("Ping server listening on port 3000...\n")
    log.Fatal(http.ListenAndServe(":3000", router))
}

