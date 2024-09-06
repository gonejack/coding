package main

import (
	"cmp"
	"fmt"
)

func insertionSort[T cmp.Ordered](arr []T) []T {
	for i, v := range arr {
		x := i - 1
		for x >= 0 && arr[x] > v {
			arr[x+1] = arr[x]
			x -= 1
		}
		arr[x+1] = v
	}
	return arr
}

func main() {
	var ss = []string{"b", "0", "01", "a", "1"}

	insertionSort(ss)

	fmt.Println(ss)
}
