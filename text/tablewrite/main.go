package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "username\tfirstname\tlastname")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich")
	fmt.Fprintln(w, "novak\tJohn\tSmith")
	w.Flush()
}
