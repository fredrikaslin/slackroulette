package roulette

import (
	"fmt"
	"net/http"
)

func run() {

	fmt.Println("staring roulette....")

	resp, err := http.Get("http://example.com/")
}
