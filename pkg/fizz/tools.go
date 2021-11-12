package fizz

import (
	"fmt"
	"log"
	"os"
)

// this "cursed" way is testing writing *directly* to the right terminal improves anything
func PrintCursed(terminal *os.File, n []byte) {
	// Will be helpful when testing TCP magic
	// 	terminal := OpenFile("/dev/stdout")
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
