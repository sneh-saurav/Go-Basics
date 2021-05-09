package main

import (
	"fmt"
)

func main() {
	var i = []int{1, 2, 3, 4, 5}
	i[2] = 6

	fmt.Printf("%v \n", i[:])
}
