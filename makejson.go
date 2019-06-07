package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	m := make(map[string]string)

	fmt.Printf("Please enter a name: ")
	fmt.Scanf("%s", &name)
	m["name"] = name

	fmt.Printf("Please enter an address:")
	fmt.Scanf("%s", &address)
	m["address"] = address

	jsonString, _ := json.Marshal(m)
	fmt.Print(string(jsonString), "\n")
}
