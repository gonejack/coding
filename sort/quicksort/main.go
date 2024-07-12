package main

import (
	"cmp"
	"fmt"
)

func partition[T cmp.Ordered](arr []T, a, b int) int {
	m := arr[b] // 选择基准值
	for a < b {
		for a < b && m >= arr[a] { // a 指针值 <= m a 指针右移
			a++
		}
		arr[b] = arr[a]            // a 指针值 > m a 值移到 b 位置
		for a < b && m <= arr[b] { // b 指针值 >= m b 指针左移
			b--
		}
		arr[a] = arr[b] // b 指针值 < m b 值移到 a 位置
	}
	arr[b] = m // m 替换 b 值
	return b
}
func quicksort[T cmp.Ordered](s []T, a, b int) {
	if b > a {
		c := partition(s, a, b) // 分区
		quicksort(s, a, c-1)    // 左边部分排序
		quicksort(s, c+1, b)    // 右边部分排序
	}
}

func main() {
	var ss = []string{"b", "2", "01", "a"}
	quicksort(ss, 0, len(ss)-1)
	fmt.Println(ss)
}
