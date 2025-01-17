package md5coder

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func Md5Up(s string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(s)))
}
