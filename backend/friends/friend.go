package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Candy struct {
	candyId int
	name    string
	pieces  int
}

func main() {
	time.Sleep(10 * time.Second)
	fmt.Println("A friend is bringing candies.")
	bringCandy()
}

func bringCandy() {
	for {
		time.Sleep(time.Duration(randInt(5, 10)) * time.Second)
		bring()
	}
}

func bring() {
	db, err := sql.Open("mysql", "gocli:init1234@tcp(db:3306)/candy")
	if err != nil {
		fmt.Println("panic...")
		panic(err.Error())
	}

	defer db.Close()

	//Update db

	updateCandy, err := db.Prepare("UPDATE candy SET pieces=pieces+? WHERE candyId=?")
	ErrorCheck(err)
	tx, er := db.Begin()
	ErrorCheck(er)
	_, e := tx.Stmt(updateCandy).Exec(randInt(50, 100), randInt(1, 3))
	ErrorCheck(e)
	commitError := tx.Commit()
	ErrorCheck(commitError)
	db.Close()

}

func randInt(mini int, maxi int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(maxi-mini+1) + mini
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
