package main

import (
    "log"
    "encoding/json"
    "net/http"
    "fmt"

    "github.com/emersonmendes/go-hw/service"
    _ "github.com/lib/pq"
)

var err error

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func getCars(w http.ResponseWriter, r *http.Request){
    var Cars = service.GetCars() 
    json.NewEncoder(w).Encode(Cars)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/cars", getCars)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    fmt.Println("Starting app ...")
    handleRequests()
}