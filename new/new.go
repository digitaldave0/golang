package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello go")
	// open input file
    fi, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
	// close fi on exit and check for its returned error
defer func() {
	if err := fi.Close(); err != nil {
		panic(err)
	}
}()
// make a buffer to keep chunks that are read
buf := make([]byte, 1024)
for {
	// read a chunk
	n, err := fi.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if n == 0 {
		break
	}

	// write a chunk
	if _, err := fo.Write(buf[:n]); err != nil {
		panic(err)
	}
}
}
