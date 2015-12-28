package main

import (
	"strconv"
	"time"
	"crypto/md5"
	"fmt"
)

func adventCoin(key string) {
	i:= 1
	for {
		incrString := strconv.Itoa(i) // converting an int to string
		keyCombin := key + incrString
		data := []byte(keyCombin)
		// had no idea I could extract values from print statements!
		br := fmt.Sprintf("%x", md5.Sum(data))
		if br[0] == '0' && br[1] == '0' && br[2] == '0' && br[3] == '0' && br[4] == '0' && br[5] == '0'{
			fmt.Println(incrString)
			fmt.Println(br)
			break
		}
		i += 1
	}
}
func main() {
	start := time.Now()
	adventCoin("yzbqklnj")
	stop := time.Since(start)
	fmt.Println("Solutions in seconds: ", stop)	
}
