package fizz

import (
	"os"
	"strconv"
)

/* const (
	NUMCHANNEL       = 10
	STRINGBUFFERSIZE = 64 // Size of hardcoded bytes value
	StepSize         = 10
	// BlockSize = STRINGBUFFERSIZE * StepSize
)
*/
const (
	BlockSize = STRINGBUFFERSIZE * STEPSIZE
)

/*
The concurrent portion of testing. Made the fixed version remove the new line thingy
*/
func FizzByteFixed(n int) [STRINGBUFFERSIZE]byte {
	// Since this bypasses string allocation, it improves performances
	// String processing in go is heavy
	DivByThree := (n%3 == 0)
	DivByFive := (n%5 == 0)
	if DivByThree && DivByFive {
		return [STRINGBUFFERSIZE]byte{70, 105, 122, 122, 98, 117, 122, 122, 10}
	} else if DivByThree {
		return [STRINGBUFFERSIZE]byte{66, 117, 122, 122, 32, 32, 32, 32, 10}
	} else if DivByFive {
		return [STRINGBUFFERSIZE]byte{70, 105, 122, 122, 32, 32, 32, 32, 10}
	} else {
		// This is good enough to convert the integer to a string
		number := strconv.Itoa(n)

		// but now use some magic to add the string to a byte array
		var output [STRINGBUFFERSIZE]byte
		counter := 0
		for i := 0; i < len(number); i++ {
			output[i] = number[i]
			counter = i

		}
		output[counter+1] = 10
		// create a new line at the end
		// output[STRINGBUFFERSIZE] = 10
		return output
	}

}

func ChunkByteFixed(n int, offset int, out chan<- *[BlockSize]byte) {
	defer wg.Done()

	// buffer to write to

	counter := 0

	var b [BlockSize]byte
	var tmp [STRINGBUFFERSIZE]byte

	// go through and calculate fizzbuzz using the byte method
	for i := n; i < n+offset; i++ {
		tmp = FizzByteFixed(i)

		// b.Write(tmp[:])
		// b +=
		for j := 0; j < STRINGBUFFERSIZE-2; j++ {
			b[counter+j] = tmp[j]
		}
		//b[counter+STRINGBUFFERSIZE-1] = 10
		counter += (STRINGBUFFERSIZE)

	}

	out <- &b
}

func ConcurrentByteFizzFixed(step int, upperlimit int) {

	var chans [NUMCHANNEL]chan *[BlockSize]byte
	for i := range chans {
		chans[i] = make(chan *[BlockSize]byte)
	}

	counter := 0
	// totalOffset := 0

	for i := 0; i < upperlimit; i += (NUMCHANNEL * step) {

		wg.Add(NUMCHANNEL)

		for j := i; j < NUMCHANNEL*step+i; j += step {
			go ChunkByteFixed(j, step, chans[counter])
			counter++
		}

		counter = 0

		for j := 0; j < NUMCHANNEL*step; j += step {
			output := <-chans[counter]
			os.Stdout.Write((*output)[:])
			counter++
		}

		counter = 0

		wg.Wait()

	}

}
