package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	n := 1000000

	fmt.Printf("n: %f \n", 1e6)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go tOn(n, wg)

	go tOn2(n, wg)

	go tOlogn(n, wg)

	wg.Wait()
}

// O(n)
func tOn(n int, wg sync.WaitGroup) {
	start1 := time.Now().UnixNano() / 1e6
	var k int64
	for i := 0; i < n; i++ {
		k++
	}
	end1 := time.Now().UnixNano() / 1e6
	fmt.Printf("n n: %d, total time %d \n", n, end1-start1)
	wg.Done()
}

// O(n^2)
func tOn2(n int, wg sync.WaitGroup) {
	start2 := time.Now().UnixNano() / 1e6
	var k int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k++
		}
	}
	end2 := time.Now().UnixNano() / 1e6
	fmt.Printf("n^2 n: %d, total time %d \n", n, end2-start2)
	wg.Done()
}

// O(nlogn)
func tOlogn(n int, wg sync.WaitGroup) {
	start3 := time.Now().UnixNano() / 1e6
	var k int64
	for i := 0; i < n; i++ {
		for j := 1; j < n; j = j * 2 {
			k++
		}
	}
	end3 := time.Now().UnixNano() / 1e6
	fmt.Printf("logn n: %d, total time %d \n", n, end3-start3)
	wg.Done()
}
