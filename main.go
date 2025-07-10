package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rishavqwerty7/BookwormApi/router"
)

func main() {
	fmt.Println("starting the server..")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000")
}
