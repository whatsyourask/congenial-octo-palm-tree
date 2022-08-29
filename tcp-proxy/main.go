package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in> ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}
func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)

	stdin_count, err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data!")
	}
	fmt.Printf("Read %d bytes from stdin\n", stdin_count)

	stdout_count, err := writer.Write(input)
	if err != nil {
		log.Fatalln("unable to write data!")
	}
	fmt.Printf("Write %d bytes from stdout\n", stdout_count)
}
