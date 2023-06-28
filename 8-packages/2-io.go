package main

import (
	"fmt"
	"os"
)

func main() {
	// first open a file
	file, err := os.Open("./data/test.txt")
	defer file.Close() // file needs to be closed whenever it exits the function

	if err != nil {
		// handle error
		return
	}

	// read the file size to know how many bytes size are needed
	stat, err := file.Stat()
	if err != nil {
		// handle error
		return
	}

	// reading the file by converting it to a Buffer struct
	bs := make([]byte, stat.Size()) // use `make()` to allow runtime array size
	_, err = file.Read(bs)          // read the Buffer into an array of bytes
	if err != nil {
		// handle error
		return
	}

	fileStr := string(bs)
	fmt.Println(fileStr)

}
