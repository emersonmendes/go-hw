package route

import (
    "encoding/json"
    "net/http"
    "github.com/emersonmendes/go-hw/service"
)

func getCars(w http.ResponseWriter, r *http.Request){
    var Cars = service.GetCars() 
    json.NewEncoder(w).Encode(Cars)
}

func CarHandleRequests() {
    http.HandleFunc("/cars", getCars)
}
