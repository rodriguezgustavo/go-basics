package main

import (
	"fmt"
	"os"
)

func main() {
	var s string

	//for initialization; condition; post { }
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println("Result 1: " + s)

	//for condition { }
	i := 1
	s = ""
	for i < len(os.Args) {
		s += os.Args[i] + " "
		i++
	}
	fmt.Println("Result 2: " + s)

	//for { }
	i = 1
	s = ""
	for {
		if i >= len(os.Args) {
			break
		}
		s += os.Args[i] + " "
		i++
	}
	fmt.Println("Result 3: " + s)

	//for _, arg := range args[x:y] { }
	s = ""
	for _, value := range os.Args[1:] {
		s += value + " "
	}
	fmt.Println("Result 4: " + s)

}
