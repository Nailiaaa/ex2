package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		randomInt := rand.Intn(100)
		numbers[i] = randomInt
	}
	fmt.Println(numbers)
}
