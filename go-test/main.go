package main

import (
	"fmt"
	"go-test/funcTest"

	"github.com/google/uuid"
)

func sayHello() {
	fmt.Print("Hi")
}

func main() {
	id := uuid.New()
	// fmt.Print("Hello world")
	sayHello()
	fmt.Println(id)
	// funcTest.SayTest()
	funcTest.SayNestedTest()

	const name string = "dddd"
	fmt.Println(name)

	mySlice := []int{1, 2, 3, 4, 5}
	mySlice = append(mySlice, 10)
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}

	myMap := make(map[string]int)

	myMap["abc"] = 5
	myMap["abcz"] = 4
	myMap["abcx"] = 6
	myMap["abcv"] = 7

	for a, b := range myMap {
		fmt.Println(a, b)
	}

	val, ok := myMap["abc"]
	if ok {
		fmt.Println("found",val)
	}


	
}
