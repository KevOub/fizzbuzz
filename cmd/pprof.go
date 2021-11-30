package main

import (
	"os"
	"runtime/pprof"

	"github.com/KevOub/fizzbuzz/pkg/fizz"
)

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	step := fizz.STEPSIZE
	up := fizz.UPPERLIMIT

	fizz.ConcurrentByteFizzFixed(step, up)

}
