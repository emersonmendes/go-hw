package main

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "database/sql"

    _ "github.com/lib/pq"
    

)

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "hw_db"
const Port = "5432"
const Password = "postgres"
const DbName = "hw_db"
const TableName = "car"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
var db *sql.DB
var err error

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Cars []Car

func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func getCars(w http.ResponseWriter, r *http.Request){

    db, err = sql.Open(PostgresDriver, DataSourceName)

    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Connected!")
    }

    defer db.Close()

    sqlStatement, err := db.Query("SELECT name, color FROM " + TableName)
    checkErr(err)

    Cars = []Car{}

    for sqlStatement.Next() {

        var car Car

        err = sqlStatement.Scan(&car.Name, &car.Color)
        checkErr(err)

        Cars = append(Cars, car)

    }

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