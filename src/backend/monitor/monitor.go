package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	candy "github.com/XGouDemo/candyStorage/src/backend/candy"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	time.Sleep(10 * time.Second)
	go handleRequest()
	foreverwaiting()
}

// myHandler implements ServeHTTP, so it is valid
type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "serving via mux-Handle")
	fmt.Println("Monitor Got Request")
}
func handleRequest() {

	// handler
	h := new(myHandler)

	// create mux and register handler
	mux := http.NewServeMux()
	mux.Handle("/", h)

	// register mux with server and listen for requests
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func foreverwaiting() {
	for {
		time.Sleep(5 * time.Second)
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
	fmt.Printf("----------------------Candy Storage--111------------------\n")
	for res.Next() {
		var candy candy.Candy
		err := res.Scan(&candy.Pieces)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("---Total candy pieces: %v\n", candy.Pieces)
	}

	//report on the most abundant candy
	res, err = db.Query("SELECT * FROM candy ORDER BY pieces DESC LIMIT 1")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var candy candy.Candy
		err = res.Scan(&candy.CandyId, &candy.Name, &candy.Pieces)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("---The most abundant candy " + candy.Name + " has " + strconv.Itoa(candy.Pieces) + " Pieces.\n")
	}
	//report on the least abundant candy
	res, err = db.Query("SELECT * FROM candy ORDER BY pieces ASC LIMIT 1")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var candy candy.Candy
		err = res.Scan(&candy.CandyId, &candy.Name, &candy.Pieces)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("---The least abundant candy " + candy.Name + " has " + strconv.Itoa(candy.Pieces) + " pieces.\n")
	}
	fmt.Printf("------------------------------------------------------\n")
}

func manyLineBreaks() {
	fmt.Printf("\n\n\n\n\n")
}
