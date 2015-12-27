package main

import (
	"fmt"
	//	"os"
	"unicode/utf8"
	"io/ioutil"
	"time"
)

func evaluate1(inputToEval string) {
	sum := 0
	sizeOfInput := utf8.RuneCountInString(inputToEval)
	for i := 0; i < sizeOfInput; i++ {
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
	sizeOfInput := utf8.RuneCountInString(inputToEval)
	for i := 0; i < sizeOfInput; i++ {
		if string(inputToEval[i]) == "(" {
			sum++
		} else if string(inputToEval[i]) == ")" {
			sum--
		}
		if sum < 0 {
			fmt.Println(i+1)
			return
			// os.Exit(0)
		}
	}
	fmt.Println(sum)
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func fromFile() {
	dat, err := ioutil.ReadFile("day1p2.txt")
	check(err)

	evaluate1(string(dat))
	evaluate2(string(dat))
}

func main() {
	start := time.Now()
	fromFile()
	stop := time.Since(start)
	fmt.Println("FileIO in seconds: ", stop)
}
