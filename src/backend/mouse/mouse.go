package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Candy struct {
	candyId int
	name    string
	pieces  int
}

func main() {
	fmt.Println("A mouse is stealing candies.")
	foreverwaiting()
}

func foreverwaiting() {
	for {
		time.Sleep(time.Duration(randInt(7, 20)) * time.Second)
		steal()
	}
}

func steal() {
	var candy Candy
	db, err := sql.Open("mysql", "gocli:init1234@tcp(db:3306)/candy")
	if err != nil {
		fmt.Println("panic...")
		panic(err.Error())
	}

	defer db.Close()

	//get the most abundant candy
	res, err := db.Query("SELECT * FROM candy ORDER BY pieces DESC LIMIT 1")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	for res.Next() {

		err = res.Scan(&candy.candyId, &candy.name, &candy.pieces)

		if err != nil {
			log.Fatal(err)
		}
	}

	//Update db

	updateCandy, err := db.Prepare("UPDATE candy SET pieces=? WHERE candyId=?")
	ErrorCheck(err)
	tx, er := db.Begin()
	ErrorCheck(er)
	var newQuantity int = candy.pieces * 7 / 10
	_, e := tx.Stmt(updateCandy).Exec(newQuantity, candy.candyId)
	ErrorCheck(e)
	commitError := tx.Commit()
	ErrorCheck(commitError)
	fmt.Println("XXXXXX-----a mouse has stolen " + strconv.Itoa(candy.pieces-newQuantity) + " pieces of " + candy.name + ".------XXXXXX")
	db.Close()

}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func randInt(mini int, maxi int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(maxi-mini+1) + mini
}
