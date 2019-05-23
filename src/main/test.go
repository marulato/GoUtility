package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/submit", func(response http.ResponseWriter, request *http.Request) {

		request.ParseForm()
		fmt.Println(request.PostForm)
		response.Write([]byte("request successfully"))
		log.Println("request Posted")
	})
	log.Println("Server started up ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
