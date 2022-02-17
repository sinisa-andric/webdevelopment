package main

import (
	"fmt"
	"net/http"
)

func handlerFunction(writter http.ResponseWriter, reader *http.Request) {

	writter.Header().Set("Content-Type", "text/html")

	// log
	fmt.Println("Someone visited oru page")

	fmt.Fprint(writter, "<h1> Welcome to my website! HOT RELOAD </h1>")

}

func main() {

	// router
	http.HandleFunc("/", handlerFunction)

	// localhost on port 3001
	http.ListenAndServe(":3001", nil)

}
