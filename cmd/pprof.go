package main

import (
	"os"
	"runtime/pprof"

	"github.com/KevOub/fizzbuzz/pkg/fizz"
)

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	step := fizz.StepSize
	fizz.ConcurrentByteFizz(step, 10000)

}
