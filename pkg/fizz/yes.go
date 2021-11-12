package fizz

import (
	"bytes"
	"os"
)

// seeing *maximum throughput* of golang by printing yes
func Yes(n int) [ChunkSize]byte {
	// Since this bypasses string allocation, it improves performances
	// String processing in go is heavy
	return [ChunkSize]byte{121, 10}
}

// Goes through iteratively with the strings being pre-compiled
func ChunkByteYes(n int, offset int) []byte {
	defer wg.Done()

	// buffer to write to
	var b bytes.Buffer
	var tmp [ChunkSize]byte

	// go through and calculate fizzbuzz using the byte method
	for i := n; i < n+offset; i++ {
		tmp = Yes(i)
		b.Write(tmp[:])

	}

	return b.Bytes()
}

func ConcurrentByteYes(step int, upperlimit int) {

	var chans [numChannels]chan []byte
	for i := range chans {
		chans[i] = make(chan []byte)
	}

	for i := 0; i < upperlimit; i += 1 {
		wg.Add(numChannels)

		for j := 0; j < numChannels*step; j += step {
			// go ChunkByteYes(j, step, chans[counter])
			go os.Stdout.Write(ChunkByteYes(j, step))

		}

		wg.Wait()

	}

}
