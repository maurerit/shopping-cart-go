package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/maurerit/shopping-cart-go/cartservice/service"
	"github.com/jinzhu/gorm"
	 _ "github.com/go-sql-driver/mysql"
)

var appname = "cartservice"

func main() {
	log.Printf("Server started")

	db, err := gorm.Open("mysql", "")

	service.DB = *db

	if err != nil {
		fmt.Println(err.Error())
	}

	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	defer service.DB.Close()
}
