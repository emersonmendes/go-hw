package database

import (
    "fmt"
    "database/sql"  
)

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "hw_db"
const Port = "5432"
const Password = "postgres"
const DbName = "hw_db"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
var db *sql.DB
var err error

func Query(query string) (*sql.Rows, error){

    db, err = sql.Open(PostgresDriver, DataSourceName)

    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Connected!")
    }

    defer db.Close()

    return db.Query(query)

}