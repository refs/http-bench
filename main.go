package main

import (
	"github.com/refs/bench/pkg/service"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe("localhost:8082", service.Dummy()); err != nil {
		log.Fatal(err)
	}
}
