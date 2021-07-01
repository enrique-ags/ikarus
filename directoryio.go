package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// main function
func main() {

	// Scan Name
	fmt.Println("Name: ")
	var first string
	fmt.Scanln(&first)
	// Scan Last Name
	fmt.Println("Enter Last Name: ")
	var second string
	fmt.Scanln(&second)

	//scan age
	fmt.Println("Age: ")
	var age int
	fmt.Scanln(&age)

	//Scan phone
	fmt.Println("Phone: ")
	var phone string
	fmt.Scanln(&phone)

	fmt.Print("Your Personal Info is: ")

	// Addition of two string
	fmt.Print(first + " " + second + " " + strconv.Itoa(age) + " " + phone)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/directory")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO directory VALUES (?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	insert.Exec(first, second, age, phone)

	defer insert.Close()

}
