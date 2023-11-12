package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Create a string reader
	reader := strings.NewReader("This is a sample stream of data.")

	// Read data from the stream using a buffer
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break // Reached the end of the stream
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
	}
}
