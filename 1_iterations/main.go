package main
import (
	"fmt"
	"os"
)
func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s :=  s + " " + os.Args[i]
		fmt.Printf("Iteration %d result: %s\n",i,s)
	}
	fmt.Println("Final result: "+s)
}


