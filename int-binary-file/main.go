package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("******** In Memory ********")
	parseInMemory("random_binary_file")
	fmt.Println("******** Unbuffered ********")
	parseUnbuffered("random_binary_file")
	fmt.Println("******** Buffered ********")
	parseBuffered("random_binary_file")
}

func parseInMemory(filename string) {
	body, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i <= len(body)-4; i = i + 4 {
		print(body[i : i+4])
	}
}

func parseUnbuffered(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	data := make([]byte, 4)
	for {
		_, err := f.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		print(data)
	}
}

func parseBuffered(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	data := make([]byte, 4)
	for {
		_, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		print(data)
	}
}

func print(data []byte) {
	// "Little Endian", least significative byte at the lower
	// memory address, closer to the end of file.
	le := uint32(data[0]) << 24
	le = le | uint32(data[1])<<16
	le = le | uint32(data[2])<<8
	le = le | uint32(data[3])

	// "Big Endian", most significative byte at the lower
	// memory address, closer to the end of a file.
	be := uint32(data[0])
	be = be | uint32(data[1])<<8
	be = be | uint32(data[2])<<16
	be = be | uint32(data[3])<<24

	fmt.Printf("%d (little endian), %d (big endian)\n", le, be)
}
