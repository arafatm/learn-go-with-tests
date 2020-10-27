package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	resp, err := http.Get(b)
	bDuration := time.Since(startB)

	fmt.Printf("\nresp = %v\n", resp)
	fmt.Printf("\nerr = %v\n", err)

	robots, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	if aDuration < bDuration {
		return a
	}

	return b
}
