package fizz

import (
	"bytes"
	"os"
	"strconv"
)

const (
	STEPSIZE         = 6000
	STRINGBUFFERSIZE = 64 // Size of hardcoded bytes value
	UPPERLIMIT       = 1000000000
	SCALER           = 100                              // This number is meant for scaling the number of iterations. This should influence the for loop
	NUMCHANNEL       = (UPPERLIMIT / SCALER) / STEPSIZE // divide the goal by the scaler then the size of each step to create art
)

/*
The concurrent portion of testing
*/

// Utilizes set []byte sizes to reduce memory allocations
func FizzByte(n int) [STRINGBUFFERSIZE]byte {
	// Since this bypasses string allocation, it improves performances
	// String processing in go is heavy
	if n%3 == 0 && n%5 == 0 {
		return [STRINGBUFFERSIZE]byte{70, 105, 122, 122, 98, 117, 122, 122, 10}
	} else if n%3 == 0 {
		return [STRINGBUFFERSIZE]byte{66, 117, 122, 122, 32, 32, 32, 32, 10}
	} else if n%5 == 0 {
		return [STRINGBUFFERSIZE]byte{70, 105, 122, 122, 32, 32, 32, 32, 10}
	} else {
		// This is good enough to convert the integer to a string
		number := strconv.Itoa(n)

		// but now use some magic to add the string to a byte array
		var output [STRINGBUFFERSIZE]byte
		for i := 0; i < len(number); i++ {
			if number[i] != 0 {
				output[i] = number[i]
			}
		}
		// create a new line at the end
		output[STRINGBUFFERSIZE-1] = 10
		return output
	}

}

// Goes through iteratively with the strings being pre-compiled
func ChunkByte(n int, offset int, out chan<- []byte) {
	defer wg.Done()

	// buffer to write to
	var b bytes.Buffer
	var tmp [STRINGBUFFERSIZE]byte

	// go through and calculate fizzbuzz using the byte method
	for i := n; i < n+offset; i++ {
		tmp = FizzByte(i)
		b.Write(tmp[:])
		// b +=

	}

	out <- b.Bytes()
}

func ConcurrentByteFizz(step int, upperlimit int) {

	var chans [NUMCHANNEL]chan []byte
	for i := range chans {
		chans[i] = make(chan []byte)
	}

	counter := 0
	// totalOffset := 0

	for i := 0; i < upperlimit; i += (NUMCHANNEL * step) {

		wg.Add(NUMCHANNEL)

		for j := i; j < NUMCHANNEL*step+i; j += step {
			go ChunkByte(j, step, chans[counter])
			counter++
		}

		counter = 0

		for j := 0; j < NUMCHANNEL*step; j += step {
			output := <-chans[counter]
			os.Stdout.Write(output[:])
			counter++
		}
		counter = 0

		wg.Wait()

	}

}
