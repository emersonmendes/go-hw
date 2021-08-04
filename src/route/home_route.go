package route

import (
    "fmt"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func HomeHandleRequests() {
    http.HandleFunc("/", homePage)
}
