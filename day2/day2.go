package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"strings"
	"regexp"
	"strconv"
)

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
	path := "day2.txt"
	dat, _ := ioutil.ReadFile(path)
	// reads input file as a string
	// leaves parsing implementation to evaluate function(s)
	evaluate1(string(dat))
	evaluate2(string(dat))
}

// line String, regexobject -> side1 side2 and side3
func reSeparate(partialInput string, re *regexp.Regexp) (a, b, c int) {
	x := re.FindAllString(partialInput, -1)
	a, _ = strconv.Atoi(x[0])
	b, _ = strconv.Atoi(x[1])
	c, _ = strconv.Atoi(x[2])
	return
}

func evaluate1(inputToEval string) {
	re := regexp.MustCompile("[0-9]+")
	lines := strings.Split(inputToEval, "\n")
	sum := 0
	for i := 0; i < len(lines)-1; i++ {
		a, b, c := reSeparate(lines[i], re)
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
		a, b, c := reSeparate(lines[i], re)
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
	fmt.Println("Rutime: ", stop)
}
