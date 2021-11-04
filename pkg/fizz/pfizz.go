package fizz

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	numChannels = 250
)

var wg sync.WaitGroup

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

func ChunkByte(n int, offset int, out chan<- []byte) {
	defer wg.Done()
	var b bytes.Buffer
	// var b [OFFSET * 9]byte
	var tmp [9]byte
	for i := n; i < n+offset; i++ {
		tmp = FizzByte(i)
		b.Write(tmp[:])
	}

	out <- b.Bytes()
}

func ConcurrentTest(step int, upperlimit int) {

	var chans [numChannels]chan string
	for i := range chans {
		chans[i] = make(chan string)
	}

	counter := 0

	for i := 0; i < upperlimit; i += numChannels {
		wg.Add(numChannels)

		for j := 0; j < numChannels*step; j += step {
			go ChunkConcurrent(j, step, chans[counter])
			counter++

		}

		counter = 0

		// fmt.Print()

		for j := i; j < numChannels+i; j++ {
			// go fmt.Print(<-chans[j])
			os.Stderr.WriteString(<-chans[counter])
			counter++
		}
		counter = 0

		wg.Wait()

	}

}

func PrintCursed(terminal *os.File, n []byte) {
	// defer terminal.Close()
	_, err := terminal.Write(n)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func OpenFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		fmt.Print(err)
		f.Close()
		return nil
	}

	return f

}

func ConcurrentTestByte(step int, upperlimit int) {
	// terminal := OpenFile("/dev/tty")

	var chans [numChannels]chan []byte
	for i := range chans {
		chans[i] = make(chan []byte)
	}

	counter := 0

	for i := 0; i < upperlimit; i += 1 {
		wg.Add(numChannels)

		for j := 0; j < numChannels*step; j += step {
			go ChunkByte(j, step, chans[counter])
			counter++
		}

		counter = 0

		// fmt.Print()

		for j := 0; j < numChannels*step; j += step {
			// go fmt.Print(<-chans[j])
			output := <-chans[counter]
			// _ = <-chans[counter]
			// output = nil
			os.Stdout.Write(output)
			// PrintCursed(terminal, output)

			counter++
		}
		counter = 0

		wg.Wait()

	}

}

func ConcurrentTestByteCursed(step int, upperlimit int) {
	terminal := OpenFile("/dev/stdout")

	var chans [numChannels]chan []byte
	for i := range chans {
		chans[i] = make(chan []byte)
	}

	counter := 0

	for i := 0; i < upperlimit; i += 1 {
		wg.Add(numChannels)

		for j := 0; j < numChannels*step; j += step {
			go ChunkByte(j, step, chans[counter])
			counter++
		}

		counter = 0

		// fmt.Print()

		for j := 0; j < numChannels*step; j += step {
			// go fmt.Print(<-chans[j])
			output := <-chans[counter]
			// _ = <-chans[counter]
			// output = nil
			// os.Stdout.Write(output)
			PrintCursed(terminal, output)

			counter++
		}
		counter = 0

		wg.Wait()

	}

}
