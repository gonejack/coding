package main

import (
	"fmt"
	"math/rand/v2"
)

func randomArr() {
	const N = 30
	var arr [N]int
	for i := range arr {
		arr[i] = i + 1
	}
	for i := range arr {
		j := i + rand.N(N-i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr[:min(N, 10)])
}

func main() {
	randomArr()
}
