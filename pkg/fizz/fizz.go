package fizz

import (
	"bytes"
	"strconv"
)

const (
	OFFSET = 256
)

func FasterItoa(n int) string {
	switch n {
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	default:
		return "0"
	}
}

func FasterItoaByte(n int) [9]byte {
	var b [9]byte
	for i := 0; i < 8; i++ {
		b[i] = (byte)(n >> i * 8)
	}

	return b
}

// Fizz(n int) returns a string depending on the mat
func Fizz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		// return strconv.FormatInt(int64(n), 10)
		return strconv.Itoa(n)
	}

}

func FizzByte(n int) [9]byte {
	if n%3 == 0 && n%5 == 0 {
		return [9]byte{70, 105, 122, 122, 98, 117, 122, 122, 10}
	} else if n%3 == 0 {
		return [9]byte{66, 117, 122, 122, 32, 32, 32, 32, 10}
	} else if n%5 == 0 {
		return [9]byte{70, 105, 122, 122, 32, 32, 32, 32, 10}
	} else {
		// return strconv.FormatInt(int64(n), 10)
		// return strconv.Itoa(n)
		number := strconv.Itoa(n)
		// output := [9]byte
		var output [9]byte
		for i := 0; i < len(number); i++ {
			if number[i] != 0 {
				output[i] = number[i]
			}
		}
		output[8] = 10
		return output
	}

}

func Chunk(n int, offset int) string {
	var b bytes.Buffer
	for i := n; i < n+offset; i++ {
		b.WriteString(Fizz(i) + "\n")
		// b.WriteString(Fizz(i))
		// b.WriteString("\n")
	}
	return b.String()
}
