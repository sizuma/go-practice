package main

import (
	"c2c/db"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello World")
	error := db.Migration(true, true)
	fmt.Println(error)
}
