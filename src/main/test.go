package main

import (
	"fmt"
	"httpGo"
	"net/http"
)

type user struct {
	Name string
	Age int
	Dick string
}

func main() {
	http.HandleFunc("/upload", func(response http.ResponseWriter, request *http.Request) {
		var myUser *user

		myUser = new(user)
		myUser.Age = 10
		myUser.Name = "大幅"
		myUser.Dick = "大鸡巴"
		fmt.Println(*myUser)
		fmt.Println(myUser)
		fmt.Println(&myUser)
		httpGo.ResponseJSON(response, myUser)

	})
	httpGo.StartServer("localhost:8080")

}
