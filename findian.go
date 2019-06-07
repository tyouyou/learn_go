package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"regexp"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Please enter a string:")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSuffix(input, "\n")
	re := regexp.MustCompile(`(?i)\Ai.*a.*n\z`)
	isMatch := re.MatchString(input)
	if isMatch {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}