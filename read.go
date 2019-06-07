package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname [20]byte
	lname [20]byte
}

func main() {
	var fileName string
	fmt.Printf("Please enter  the name of the text file: ")
	fmt.Scanf("%s", &fileName)

	f, _ := os.Open(fileName)

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	var slic []name

	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		n := name{}
		copy(n.fname[:], words[0])
		copy(n.lname[:], words[1])
		slic = append(slic, n)
	}
	fmt.Printf("%s", slic)
	fmt.Printf("\n")

	f.Close()

}
