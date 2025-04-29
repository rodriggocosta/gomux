package main

import (
	"apigo/router"
	"log"
	"net/http"
)

func main() {

	mux := router.NewRouter()

	log.Fatal(http.ListenAndServe(":9999", mux))
}
