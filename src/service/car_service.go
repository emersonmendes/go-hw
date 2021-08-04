package service

import (
    "github.com/emersonmendes/go-hw/domain"
    "github.com/emersonmendes/go-hw/database"
    _ "github.com/lib/pq"
)

func GetCars() []domain.Car {

    sqlStatement, err := database.Query("select name, color from car")
  
    checkErr(err)

    var Cars = []domain.Car{}

    for sqlStatement.Next() {

        var car domain.Car

        err = sqlStatement.Scan(&car.Name, &car.Color)
        checkErr(err)

        Cars = append(Cars, car)

    }

    return Cars

}

func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}
