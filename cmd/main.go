package main

import (
	"github.com/KevOub/fizzbuzz/pkg/fizz"
)

// const (
// 	numChannels = 500
// )

// var wg sync.WaitGroup

// func ChunkConcurrent(n int, offset int, ch chan<- string) {
// 	defer wg.Done()
// 	var b bytes.Buffer
// 	for i := n; i < n+offset; i += 5 {
// 		b.WriteString(fizz.Fizz(i) + "\n")
// 	}
// 	ch <- b.String()
// }

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()
	// fizz.ConcurrentTest(2048, 10000)
	// fmt.Print(fizz.FasterItoaByte(9))
	// fizz.ConcurrentTest(2048, 10000)

	// fmt.Print([]byte("\n"))

	// logging functions
	// f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	//pause for debugger

	// log.SetOutput(f)
	step := fizz.STEPSIZE
	up := fizz.UPPERLIMIT
	// up := fizz.UPPERLIMIT

	// fizz.ConcurrentByteFizz(15000, 100000)
	// fizz.ConcurrentByteFizz(step, 10000)
	fizz.ConcurrentByteFizzFixed(step, up)
	// timeTrack(time.Now(), "a:")

	// fizz.ConcurrentTestByteCursed(7500*4, 1000000)
	// fizz.ConcurrentTestByteCursed(100, 1000000)
	// timeTrack(time.Now(), "b:")

}
