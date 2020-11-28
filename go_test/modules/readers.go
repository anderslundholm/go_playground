package modules

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ReadFile is a collection of examples of opening up and reading a local file in different ways.
func ReadFile(filePath string) {
	// Basic file reading, whole file saved in memory
	data, err := ioutil.ReadFile(filePath)
	checkErr(err)
	fmt.Printf("ioutil.ReadFile: %s\n", string(data))

	// Open file to get an os.File value
	f, err := os.Open(filePath)
	checkErr(err)

	// Close file when done.
	defer f.Close()

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkErr(err)
	fmt.Printf("os.Open - File.Read %d bytes: %s\n", n1, string(b1[:n1]))

	// Seek to a known location in the file and Read from there.
	o2, err := f.Seek(2, 0)
	checkErr(err)
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	checkErr(err)
	fmt.Printf("os.Open - File.Seek %d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// More robustly implementation of above read with io.ReadAtLeast.
	o3, err := f.Seek(2, 0)
	checkErr(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	checkErr(err)
	fmt.Printf("os.Open - os.ReadAtLeast %d bytes @ %d: %s\n", n3, o3, string(b3))

	// Rewind with Seek(0, 0).
	_, err = f.Seek(0, 0)
	checkErr(err)

	// The buffered reader in bufio package offers additional reading methods and it's efficient with many small reads.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	checkErr(err)
	fmt.Printf("os.Open - bufio - Reader.Peek 5 bytes: %s\n", string(b4))
}
