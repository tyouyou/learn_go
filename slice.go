package main
import (
	"fmt"
	"strconv"
	"sort"
)

func main () {
	var slic = make([]int, 3)
	slic = slic[0:0]
	for {
		var input string
		fmt.Printf("Please enter an integer to be added to the slice: ")
		fmt.Scanf("%s", &input)
		if input == "X" {
			break
		} else {
			number, _ := strconv.Atoi(input)
			slic = append(slic, number)	
			sort.Sort(sort.IntSlice(slic))
			fmt.Println("The slice is: ", slic)	
			
		}
	}

}