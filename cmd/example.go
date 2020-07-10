package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/c2h5oh/hide"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Random IDs")
	for i := 0; i < 10; i++ {
		v := rand.Int31n(1000000)
		o := hide.Int32Obfuscate(v, nil)

		fmt.Printf("%8d -> %10d\n", v, o)
	}

	fmt.Println("\nConsecutive IDs")
	start := rand.Int31n(1000000)
	for i := start; i < start+10; i++ {
		o := hide.Int32Obfuscate(i, nil)

		fmt.Printf("%8d -> %10d\n", i, o)
	}
}
