package main

import (
	"fmt"
	"net/http"
)

func handlerFunction(writter http.ResponseWriter, reader *http.Request) {

	writter.Header().Set("Content-Type", "text/html")

	// log
	fmt.Println("Someone visited oru page")
	fmt.Println(writter, reader.URL.Path)

	if reader.URL.Path == "/" {
		fmt.Fprint(writter, "<h1> Welcome to my website! </h1>")
	} else if reader.URL.Path == "/contact" {
		fmt.Fprint(writter, "To get in touch, please send an email to <a href=\"mailto:sinisa.andric@mail.ru\"> sinisa.andric@mail.ru")
	}
}

func main() {

	// router
	http.HandleFunc("/", handlerFunction)

	// localhost on port 3001
	http.ListenAndServe(":3001", nil)

}
