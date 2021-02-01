package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	createDocRoot()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func createDocRoot() {
	err := os.Mkdir("./static", 0755)
	if err != nil {
		log.Fatal(err)
	}
}
