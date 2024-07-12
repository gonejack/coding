package main

import "C"
import (
	"fmt"
	"sync"
)

//export Add
func Add(a, b int) int {
	var g sync.WaitGroup

	g.Add(1)
	go func() {
		a += 1
		g.Done()
	}()

	g.Add(1)
	go func() {
		b += 1
		g.Done()
	}()

	g.Wait()

	return a + b + 1
}

//export Fmt
func Fmt(s *C.char) string {
	return fmt.Sprintf("formated %s", C.GoString(s))
}

func main() {
	// build .a file
	// go build -v -buildmode=c-shared libx.go

	// build .so file
	// go build -v -buildmode=c-archive libx.go
}
