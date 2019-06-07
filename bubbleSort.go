package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 10 integer in the followings")
	var inputInts []int
	for i := 1; i <= 10; i++ {
		fmt.Printf("-> No.%d : ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		inputNum, _ := strconv.Atoi(input)
		inputInts = append(inputInts, inputNum)
	}
	bubbleSort(inputInts)
	fmt.Println(inputInts, "\n")
}

func bubbleSort(slices []int) {
	for i := 0; i < len(slices)-1; i++ {
		for j := 0; j < len(slices)-i-1; j++ {
			if slices[j] > slices[j+1] {
				swap(slices, j)
			}
		}
	}
}

func swap(slices []int, i int) {
	temp := slices[i]
	slices[i] = slices[i+1]
	slices[i+1] = temp
}
