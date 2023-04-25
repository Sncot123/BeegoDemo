package main

import "fmt"

type TestStruct struct {
	Id   int
	Name string
}

func main() {
	var test1 = TestStruct{Id: 18, Name: "asda"}
	fmt.Print(test1)
}
