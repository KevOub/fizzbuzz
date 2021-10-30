package Fizz

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < 10000; i++ {
		Fizz(i)
	}
}
