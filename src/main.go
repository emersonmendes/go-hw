package main

import (
    "net/http"
    "fmt"
    "github.com/emersonmendes/go-hw/route"
)

func buildRoutes() {
    route.CarHandleRequests()
    route.HomeHandleRequests()
}

func main() {
    buildRoutes()
    fmt.Println("Starting app ...")
    http.ListenAndServe(":10000", nil)
}