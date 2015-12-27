package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// simplifies text reading operation
func fromFile() {
	path := "day3.txt"
	dat, err := ioutil.ReadFile(path)
	check(err)
	evaluate1(string(dat))
	evaluate2(string(dat))
	
	//	walking before running
	// test := "^v"
	// fmt.Println(len(test))
	// evaluate2(test)
}

// simple two-dimensional structure
type coord struct {
	x int
	y int
}

// *coord, *[]coord => bool
func visited(entity *coord, coordinateSlice *[]coord) bool{
	state := true
	for _, v := range(*coordinateSlice) {
		if v == *entity {
			state = false
		}
	}
	return state
}
			
func motion(a string, entity *coord, locationsVisited *int, coordinateSlice *[]coord) {
	if a == "^" {
		entity.y++
	} else if a == "<" {
		entity.x--
	} else if a == ">" {
		entity.x++
	} else if a == "v" {
		entity.y--
	}
	if visited(entity, coordinateSlice) {
		fmt.Println(*entity)
		*coordinateSlice = append(*coordinateSlice, *entity)
		*locationsVisited++
	}
}

func evaluate1(inputToEval string) {
	start := coord{x: 20, y: 20}
	var listOfCoords []coord
	listOfCoords = append(listOfCoords, start)
	fmt.Println(start)
	numberOfHouses := 1
	fmt.Println(inputToEval)
	for i := 0; i < len(inputToEval); i++ {
		motion(string(inputToEval[i]), &start, &numberOfHouses, &listOfCoords)
	}
	fmt.Println(numberOfHouses)
}

func evaluate2(inputToEval string) {
	start := coord{x: 20, y: 20}
	var listOfCoords []coord
	listOfCoords = append(listOfCoords, start)
	fmt.Println(start)
	// now two houses because s and rs start at the same place
	numberOfHouses := 1
	fmt.Println(inputToEval)
	for i := 0; i < len(inputToEval); i+=2 {
		motion(string(inputToEval[i]), &start, &numberOfHouses, &listOfCoords)
	}
	// resetting coordinates for rs
	start = coord{x: 20, y: 20}
	for i := 1; i < len(inputToEval); i+=2 {
		motion(string(inputToEval[i]), &start, &numberOfHouses, &listOfCoords)
	}
	fmt.Println(numberOfHouses)
}

func main() {
	start := time.Now()
	fromFile()
	stop := time.Since(start)
	fmt.Println("Solutions in seconds: ", stop)
}
