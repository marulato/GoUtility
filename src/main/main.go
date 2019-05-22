package main

import (
	arrayList "arraylist"
	"fmt"
)

func main() {
	mylist := arrayList.New()
	mylist.Add("hello")
	mylist.Add("Golang")

	secList := arrayList.New()
	secList.Add("hello")
	secList.Add("java")

	mylist.AddAll(secList)
	fmt.Println(mylist.Size())
	for i := 0; i < mylist.Size(); i++ {
		fmt.Println(mylist.Get(i))
	}
}
