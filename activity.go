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
	fmt.Print("Please enter value for acceleration: ")
	acce, _ := reader.ReadString('\n')
	acce = strings.TrimSuffix(acce, "\n")
	acceNum, _ := strconv.ParseFloat(acce, 64)

	fmt.Print("Please enter value for initial velocity: ")
	velo, _ := reader.ReadString('\n')
	velo = strings.TrimSuffix(velo, "\n")
	veloNum, _ := strconv.ParseFloat(velo, 64)

	fmt.Print("Please enter value for initial displacement: ")
	disp, _ := reader.ReadString('\n')
	disp = strings.TrimSuffix(disp, "\n")
	dispNum, _ := strconv.ParseFloat(disp, 64)

	fmt.Print("Please enter a value for time: ")
	time, _ := reader.ReadString('\n')
	time = strings.TrimSuffix(time, "\n")
	timeNum, _ := strconv.ParseFloat(time, 64)

	dispFunc := genDispalceFn(acceNum, veloNum, dispNum)
	dispRes := dispFunc(timeNum)

	fmt.Println("Computing the displacment.... ")
	fmt.Printf("The result is: %f \n", dispRes)
}

func genDispalceFn(acce float64, velo float64, initDispl float64) func(t float64) float64 {
	return func(t float64) float64 {
		disp := (0.5 * acce * (t * t)) + (velo * t) + initDispl
		return disp
	}
}
