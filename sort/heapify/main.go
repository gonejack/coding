package main

import "fmt"

func main() {
	var arr = []int{0, 1, 2, 4, 6}

	buildHeap(arr, len(arr))

	fmt.Println(arr)
}

func heapify(arr []int, p, n int) {
	x := p
	l := p*2 + 1
	r := p*2 + 2
	if l < n && arr[l] > arr[x] { // 左节点是否大于父节点
		x = l
	}
	if r < n && arr[r] > arr[x] { // 右节点是否大于父节点
		x = r
	}
	if p != x {
		arr[x], arr[p] = arr[p], arr[x]
		heapify(arr, x, n)
	}
}
func buildHeap(arr []int, len int) {
	p := len/2 - 1 // 找到最后一个节点的父节点
	for p >= 0 {
		heapify(arr, p, len)
		p--
	}
}
