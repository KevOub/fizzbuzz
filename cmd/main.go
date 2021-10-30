package main

import (
	"bytes"
	"fmt"
)

func Chunk(n int, offset int) string {
	var b bytes.Buffer
	for i := n; i < n+offset; i++ {
		b.WriteString(Fizz(i))
		b.WriteString("\n")
		// fmt.Println(Fizz(i))
	}
	return b.String()
}

func main() {
	step := 255

	for i := step; i < 1000000000; i += step {
		fmt.Print(Chunk(i+step, step))
	}
}
