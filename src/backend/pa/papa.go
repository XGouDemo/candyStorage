package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	for {
		time.Sleep(3 * time.Second)
		//url := "http://172.21.0.3:8080"
		url := "http://candystorage_monitor_1:8080"
		fmt.Println("Papa is asking: " + url)
		resp, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		fmt.Println(resp)

	}

}
