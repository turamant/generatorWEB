package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", homePageHandler)
	fmt.Println("Server listening on port 4000")
	log.Panic(
		http.ListenAndServe(":4000", nil),
	)

}


func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}