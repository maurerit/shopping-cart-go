package main

import (
	"github.com/maurerit/shopping-cart-go/sw"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
