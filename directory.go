package main

import (

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/directory")
    
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    insert, err := db.Query("INSERT INTO directory VALUES ('Sam','Smith',4,'xxx-xxx-xxxx'),('Elliot','Gutierrez',55,'yyy-ddd-9999')")
    
    if err != nil {
        panic(err.Error())
    }

    defer insert.Close()

}