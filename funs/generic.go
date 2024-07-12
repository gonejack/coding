package funs

func Empty[T comparable](v T) bool {
	var zero T
	return v == zero
}
