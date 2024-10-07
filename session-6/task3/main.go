package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	slc := make([]int, 3, 5)

	for {
		if len(slc) > 5 {
			break
		}
		fmt.Println("Slice elements:", slc)
		fmt.Println("Length of slice:", len(slc))
		fmt.Println("Capacity of slice:", cap(slc))
		slc = append(slc, rand.IntN(math.MaxInt))

	}

}
