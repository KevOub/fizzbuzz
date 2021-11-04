package Fizz_testing

import (
	"testing"

	"github.com/KevOub/fizzbuzz/pkg/fizz"
)

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizz.ConcurrentTest(1024, 10000)
	}
}
