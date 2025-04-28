package main

import (
	"apigo/infra"
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	dbConnect, err := infra.Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println(dbConnect)

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	})

	log.Fatal(http.ListenAndServe(":9999", mux))
}
