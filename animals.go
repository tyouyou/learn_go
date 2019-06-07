package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please type your animal and your request about the animal > ")
		request, _ := reader.ReadString('\n')
		request = strings.TrimSuffix(request, "\n")
		requests := strings.Fields(request)
		res := requests[1]
		switch animal := requests[0]; animal {
		case "cow":
			reflect.ValueOf(cow).MethodByName(res).Call([]reflect.Value{})
		case "bird":
			reflect.ValueOf(bird).MethodByName(res).Call([]reflect.Value{})
		case "snake":
			reflect.ValueOf(snake).MethodByName(res).Call([]reflect.Value{})
		}
	}
}
