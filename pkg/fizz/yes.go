package fizz

import (
	"os"
)

/* const (
	NUMCHANNEL       = 10
	STRINGBUFFERSIZE = 64 // Size of hardcoded bytes value
	StepSize         = 10
	// BlockSize = STRINGBUFFERSIZE * StepSize
)
*/
/*
The concurrent portion of testing. Made the fixed version remove the new line thingy
*/
func YesByteFixed(n int) [STRINGBUFFERSIZE]byte {
	// Since this bypasses string allocation, it improves performances
	// String processing in go is heavy
	return [STRINGBUFFERSIZE]byte{121, 10}

}

func ChunkByteFixedYes(n int, offset int, out chan<- *[BlockSize]byte) {
	defer wg.Done()

	// buffer to write to

	counter := 0

	var b [BlockSize]byte
	var tmp [STRINGBUFFERSIZE]byte

	// go through and calculate fizzbuzz using the byte method
	for i := n; i < n+offset; i++ {
		tmp = YesByteFixed(i)

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

func ConcurrentByteYesFixed(step int, upperlimit int) {

	var chans [NUMCHANNEL]chan *[BlockSize]byte
	for i := range chans {
		chans[i] = make(chan *[BlockSize]byte)
	}

	counter := 0
	// totalOffset := 0

	for i := 0; i < upperlimit; i += (NUMCHANNEL * step) {

		wg.Add(NUMCHANNEL)

		for j := i; j < NUMCHANNEL*step+i; j += step {
			go ChunkByteFixedYes(j, step, chans[counter])
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
