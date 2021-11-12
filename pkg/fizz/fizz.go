package fizz

import (
	"bytes"
	"strconv"
)

const (
	OFFSET = 256
)

// Fizz(n int) returns a string depending on the mat
func Fizz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		// return strconv.FormatInt(int64(n), 10)
		return strconv.Itoa(n)
	}

}

// Iterative approach to fizz buzz with chunking
func Chunk(n int, offset int) string {
	var b bytes.Buffer
	for i := n; i < n+offset; i++ {
		b.WriteString(Fizz(i) + "\n")
		// b.WriteString(Fizz(i))
		// b.WriteString("\n")
	}
	return b.String()
}
