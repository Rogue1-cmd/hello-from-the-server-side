package main

import (
	"fmt"
	"log"
	"net/http"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	fileServer := http.FileServer(http.Dir("./files"))
	http.Handle("/", fileServer)

	fmt.Println("Server is Starting at port 8080")
	// http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
