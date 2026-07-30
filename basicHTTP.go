package main

import (
	"fmt"      //standard i/o
	"log"      // logging messages to console
	"net/http" //user for building http server and clients
)

const portNum string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "info page \n")
	fmt.Fprintf(w, "hi world")
}

func main() {

	log.Println("startng our simple http server.")
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	log.Println("started on port", portNum)
	fmt.Println("To close connect do Ctrl + C")

	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
