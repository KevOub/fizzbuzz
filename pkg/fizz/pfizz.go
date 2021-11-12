package fizz

import (
	"strings"
	"sync"
)

var wg sync.WaitGroup

// Innefficient chunking. String allocation is taxing on a computer
func ChunkConcurrent(n int, offset int, ch chan<- string) {
	defer wg.Done()
	// var b bytes.Buffer
	// var t string - BAD
	var builder strings.Builder

	for i := n; i < n+offset; i++ {
		// b.WriteString(Fizz(i) + "\n")
		// t += Fizz(i) + "\n"
		builder.WriteString(Fizz(i))
		builder.WriteString("\n")
	}
	ch <- builder.String()
}
