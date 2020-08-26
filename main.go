package main

import (
	"fmt"
	"log"
	"net/http"

	"./routes"
)

func main() {

	fmt.Println("API Started.")

	router := routes.GetRouter()
	port := ":3000"

	err := http.ListenAndServe(port, router)
	log.Fatal(err)

}
