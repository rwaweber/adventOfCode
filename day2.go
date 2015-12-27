package main

import (
	"fmt"
	//	"unicode/utf8"
	"io/ioutil"
	"time"
	"strings"
	"regexp"
	"strconv"
)

// necessary for string converstion
// handles any conversion errors with panic()
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// one-line less than comparison
func smallest(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// simplifies text reading operation
func fromFile() {
	path := "day2p1.txt"
	dat, err := ioutil.ReadFile(path)
	check(err)
	// reads input file as a string
	// leaves parsing implementation to evaluate function(s)
	evaluate1(string(dat))
	evaluate2(string(dat))
}

func evaluate1(inputToEval string) {
	lines := strings.Split(inputToEval, "\n")
	sum := 0
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(lines)-1; i++ {
		x := re.FindAllString(lines[i], -1)
		a, err := strconv.Atoi(x[0])
		check(err)
		b, err := strconv.Atoi(x[1])
		check(err)
		c, err := strconv.Atoi(x[2])
 		check(err)
		side1, side2, side3 := a*b, a*c, b*c
		sum += (2*side1)+(2*side2)+(2*side3)
		sum += smallest(side1, smallest(side2, side3))
	}
	fmt.Println(sum)
}

func evaluate2(inputToEval string) {
	lines := strings.Split(inputToEval, "\n")
	sum := 0
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(lines)-1; i++ {
		x := re.FindAllString(lines[i], -1)
		a, err := strconv.Atoi(x[0])
		check(err)
		b, err := strconv.Atoi(x[1])
		check(err)
		c, err := strconv.Atoi(x[2])
		check(err)
		// important distinction is that the dimensions are added
		// instead of multiplied
		face1, face2, face3 := a+b, a+c, b+c		
		k := smallest(face1, smallest(face2, face3))
		sum += (a*b*c) + 2*k
	}
	fmt.Println(sum)
}

func main() {
	start := time.Now()
	fromFile()
	stop := time.Since(start)
	fmt.Println("Both solutions in seconds: ", stop)
}
