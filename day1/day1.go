package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func evaluate1(inputToEval string) {
	sum := 0
	for i := 0; i < len(inputToEval); i++ {
		if string(inputToEval[i]) == "(" {
			sum++
		} else if string(inputToEval[i]) == ")" {
			sum--
		}
	}
	fmt.Println(sum)
}

func evaluate2(inputToEval string) {
	sum := 0
	for i := 0; i < len(inputToEval); i++ {
		if string(inputToEval[i]) == "(" {
			sum++
		} else if string(inputToEval[i]) == ")" {
			sum--
		}
		if sum < 0 {
			fmt.Println(i+1)
			return
		}
	}
	fmt.Println(sum)
}

func fromFile() {
	dat, _ := ioutil.ReadFile("day1.txt")
	evaluate1(string(dat))
	evaluate2(string(dat))
}

func main() {
	start := time.Now()
	fromFile()
	stop := time.Since(start)
	fmt.Println("Runtime: ", stop)
}
