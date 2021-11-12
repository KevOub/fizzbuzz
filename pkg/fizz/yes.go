package fizz

// seeing *maximum throughput* of golang by printing yes
func Yes(n int) [ChunkSize]byte {
	// Since this bypasses string allocation, it improves performances
	// String processing in go is heavy
	return [ChunkSize]byte{121, 10}
}
