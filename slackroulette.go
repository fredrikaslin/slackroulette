package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("staring roulette....")

	resp, err := http.Get("http://example.com/")
}
