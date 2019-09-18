package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	initDb()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":1111", router))
}
