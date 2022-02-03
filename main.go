package main

import (
	"fmt"
	"net/http"
)

func handlerFunction(writter http.ResponseWriter, reader *http.Request) {

	fmt.Fprint(writter, "<h1> Welcome to my website! </h1>")

}

func main() {

	// router
	http.HandleFunc("/", handlerFunction)

	// localhost on port 3001
	http.ListenAndServe(":3001", nil)

}
