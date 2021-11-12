package fizz

import (
	"bytes"
	"os"
	"strconv"
)

const (
	numChannels = 1300
	ChunkSize   = 64
)

/*
The concurrent portion of testing
*/

// Utilizes set []byte sizes to reduce memory allocations
func FizzByte(n int) [ChunkSize]byte {
	// Since this bypasses string allocation, it improves performances
	// String processing in go is heavy
	if n%3 == 0 && n%5 == 0 {
		return [ChunkSize]byte{70, 105, 122, 122, 98, 117, 122, 122, 10}
	} else if n%3 == 0 {
		return [ChunkSize]byte{66, 117, 122, 122, 32, 32, 32, 32, 10}
	} else if n%5 == 0 {
		return [ChunkSize]byte{70, 105, 122, 122, 32, 32, 32, 32, 10}
	} else {
		// This is good enough to convert the integer to a string
		number := strconv.Itoa(n)

		// but now use some magic to add the string to a byte array
		var output [ChunkSize]byte
		for i := 0; i < len(number); i++ {
			if number[i] != 0 {
				output[i] = number[i]
			}
		}
		// create a new line at the end
		output[ChunkSize-1] = 10
		return output
	}

}

// Goes through iteratively with the strings being pre-compiled
func ChunkByte(n int, offset int, out chan<- []byte) {
	defer wg.Done()

	// buffer to write to
	var b bytes.Buffer
	var tmp [ChunkSize]byte

	// go through and calculate fizzbuzz using the byte method
	for i := n; i < n+offset; i++ {
		tmp = FizzByte(i)
		b.Write(tmp[:])

	}

	out <- b.Bytes()
}

func ConcurrentByteFizz(step int, upperlimit int) {

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

		for j := 0; j < numChannels*step; j += step {
			output := <-chans[counter]
			os.Stdout.Write(output[:])
			counter++
		}
		counter = 0

		wg.Wait()

	}

}
