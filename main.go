package main

import (
	"fmt"
	"time"
	"math/rand"
)

func binarySearch(n int, t int, arr ...int) {
	if n < 1 || n > t {
		fmt.Println("n must be > 0 & <", t)
		return
	}
	op := 0
	for {
		op++
		m := len(arr)/2
		arrV := arr[m]
		if n == arrV {
			fmt.Printf("Found %d, in %d operations.", n, op)
			return
		} else if n < arrV {
			fmt.Println("<", arrV)
			arr = arr[:m]
		} else if n > arrV {
			fmt.Println(">", arrV)
			arr = arr[m:]
		}
	}
	fmt.Println(7)
}

func main() {
	const t = 1000000
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(t)
	arr := []int{}
	for i := 0; i <= t; i++ {
		arr = append(arr, i)
	}
	binarySearch(r, t, arr...)
}
