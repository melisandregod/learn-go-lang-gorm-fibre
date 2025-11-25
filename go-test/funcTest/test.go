package funcTest

import (
	"fmt"
	functestin "go-test/funcTest/internal/funcTestin"
)

func SayNestedTest() {
	fmt.Println("hi momo")
	SayTest()
	functestin.SayInternal()
}
