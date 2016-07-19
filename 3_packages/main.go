package main

import (
	"fmt"
	"github.com/rodriguezgustavo/go-basics/3_packages/math"
	"os"
	"strconv"
)

func main() {
	var result int64
	for _, val := range os.Args[1:] {
		value, err := strconv.ParseInt(val, 10, 0)
		if err == nil {
			result = math.Add(result, value)
		}
	}
	fmt.Printf("Suma: %d\n", result)
}
