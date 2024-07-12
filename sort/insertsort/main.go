package main

import (
	"cmp"
	"fmt"
)

func insertionSort[T cmp.Ordered](arr []T) []T {
	for i, v := range arr {
		p := i - 1
		for p >= 0 && arr[p] > v {
			arr[p+1] = arr[p]
			p -= 1
		}
		arr[p+1] = v
	}
	return arr
}

func main() {
	var ss = []string{"b", "0", "01", "a", "1"}

	insertionSort(ss)

	fmt.Println(ss)
}
