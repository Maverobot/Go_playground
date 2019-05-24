package main

import (
	"log"

	"github.com/imroc/req"
)

func main() {
	r, err := req.Get("http://localhost:8081/articles")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", r.Dump())
}
