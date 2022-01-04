package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	a := make(chan int)
	go func(a chan int) {
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 2
			a <- 1
			<-a
		}
	}(a)

	go func(a chan int) {
		wg.Done()
		i := 2
		for i <= 10 {
			<-a
			fmt.Println(i)
			i = i + 2
			a <- 1
		}
	}(a)

	o := runtime.NumGoroutine()
	fmt.Printf("NumGoroutine %d", o)
	// time.Sleep(time.Second)
	wg.Wait()
}

func printO(a chan int) {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 2
		a <- 1
		<-a
	}
}

func printS(a chan int, wg sync.WaitGroup) {
	wg.Done()
	i := 2
	for i <= 10 {
		<-a
		fmt.Println(i)
		i = i + 2
		a <- 1
	}
}
