package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started")

	log.Fatal(http.ListenAndServe(":6000", r))
	fmt.Println("Listening at port 6000 ...")
}
