package main

import (
	"fmt"
	"os"
)

func main() {
	title := os.Args[1]
	err := generateFile("./", title)
	if err != nil {
		fmt.Printf("leetgen: %s\n", err.Error())
	}
}
