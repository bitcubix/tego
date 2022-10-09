package main

import (
	"fmt"

	"{{mod}}/test1"
	"{{mod}}/test2"
)

func main() {
	fmt.Println("example template")
	fmt.Println(test1.Test())
	fmt.Println(test2.Test())
}
