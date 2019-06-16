package main

import (
	"net/http"

	"github.com/riyadennis/goophercises/testing/signal"
)

func main() {
	http.HandleFunc("/handle", signal.Handler)
	http.ListenAndServe(":8080", nil)
}
