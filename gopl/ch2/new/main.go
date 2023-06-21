package main

import "fmt"

func main() {
	fmt.Println(delta(6, 3))
}
func delta(old, new int) int {
	return new - old
}
