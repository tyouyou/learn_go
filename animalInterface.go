package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name      string
	food      string
	locomtion string
	noise     string
}

type Bird struct {
	name      string
	food      string
	locomtion string
	noise     string
}

type Snake struct {
	name      string
	food      string
	locomtion string
	noise     string
}

func (c Cow) Eat()     { fmt.Println(c.food) }
func (c Cow) Move()    { fmt.Println(c.locomtion) }
func (c Cow) Speak()   { fmt.Println(c.noise) }
func (b Bird) Eat()    { fmt.Println(b.food) }
func (b Bird) Move()   { fmt.Println(b.locomtion) }
func (b Bird) Speak()  { fmt.Println(b.noise) }
func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomtion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func CreatAnimal(newAnimalName string, newAnimalType string) (Animal, error) {
	switch newAnimalType {
	case "cow":
		return Cow{
			name:      newAnimalName,
			food:      "grass",
			locomtion: "walk",
			noise:     "moo",
		}, nil
	case "snake":
		return Snake{
			name:      newAnimalName,
			food:      "mice",
			locomtion: "slither",
			noise:     "hsss",
		}, nil
	case "bird":
		return Bird{
			name:      newAnimalName,
			food:      "sworms",
			locomtion: "fly",
			noise:     "peep",
		}, nil
	}
	return nil, errors.New("Failed! Please type correct Animal Type(cow, snake, bird)!")
}

func main() {
	animals := make(map[string]Animal)
	fmt.Println("Welcome to Animal Application!")
	fmt.Println("It's a infinite loop. If you want to exit. Please press Ctrl+C")
	fmt.Println("Input Format of creat new animal: newanimal <animal name> <animal type>")
	fmt.Println("Input Format of query about animal: query <animal name> <request>")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		request, _ := reader.ReadString('\n')
		request = strings.TrimSuffix(request, "\n")
		requests := strings.Fields(request)
		req := requests[0]
		length := len(requests)
		if length != 3 {
			fmt.Println("Please type with the correct format!")
			continue
		}
		switch req {
		case "newanimal":
			newAnimalName := requests[1]
			newAnimalType := requests[2]
			newAnimal, err := CreatAnimal(newAnimalName, newAnimalType)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Created it!")
			animals[newAnimalName] = newAnimal
		case "query":
			animalName := requests[1]
			activity := requests[2]
			animal, ok := animals[animalName]
			if !ok {
				fmt.Println("Failed! The animal name do not exist!")
				break
			}
			switch activity {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Failed! Please type eat, move or speak in the third string!")
			}
		default:
			fmt.Println("Failed! Please type newanimal or query in the first string!")
		}
	}

}
