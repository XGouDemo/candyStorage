package main

import (
	"database/sql"
	"fmt"
	"log"
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
	fmt.Println("monitor reporting...")
	foreverwaiting()
}

func foreverwaiting() {
	sum := 0
	for {
		sum++ // repeated forever
		time.Sleep(3 * time.Second)
		reportCandyStorage()
	}
}

func reportCandyStorage() {
	manyLineBreaks()
	db, err := sql.Open("mysql", "gocli:init1234@tcp(db:3306)/candy")
	if err != nil {
		fmt.Println("panic...")
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT SUM(pieces) FROM candy")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
	fmt.Printf("----------------------Candy Storage----------------------\n")
	for res.Next() {
		var candy Candy
		err := res.Scan(&candy.pieces)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("---Total candy pieces: %v\n", candy.pieces)
	}

	//report on the most abundant candy
	res, err = db.Query("SELECT * FROM candy ORDER BY pieces DESC LIMIT 1")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var candy Candy
		err = res.Scan(&candy.candyId, &candy.name, &candy.pieces)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("---The most abundant candy " + candy.name + " has " + strconv.Itoa(candy.pieces) + " pieces.\n")
	}
	//report on the least abundant candy
	res, err = db.Query("SELECT * FROM candy ORDER BY pieces ASC LIMIT 1")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var candy Candy
		err = res.Scan(&candy.candyId, &candy.name, &candy.pieces)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("---The least abundant candy " + candy.name + " has " + strconv.Itoa(candy.pieces) + " pieces.\n")
	}
	fmt.Printf("------------------------------------------------------\n")
}

func manyLineBreaks() {
	fmt.Printf("\n\n\n\n\n")
}
