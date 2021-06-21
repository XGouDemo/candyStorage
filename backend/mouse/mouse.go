package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type candy struct {
	candyId int
	name    string
	pieces  int
}

func main() {
	fmt.Println("A mouse is stealing candies.")
	foreverwaiting()
}

func foreverwaiting() {
	sum := 0
	for {
		sum++ // repeated forever
		time.Sleep(3 * time.Second)
		openDBConn()
	}
}

func openDBConn() {
	db, err := sql.Open("mysql", "gocli:init1234@tcp(db:3306)/candy")
	if err != nil {
		fmt.Println("panic...")
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("<<<<<<<<<<<<<Ping failed...")
		fmt.Println(err)
	} else {
		fmt.Println("Mouse PING SUCCESS.")
	}
}
