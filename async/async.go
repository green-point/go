
package main

import (
	"time"
	"sync"
	"fmt"
)

func say_hello() {
	fmt.Println("Hello, World")
}

func say_sample() {
	go say_hello()
	fmt.Println("Hello again")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func sum_sample() {
	s := []int {1, 2, 3, 4, 5, 6}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x := <- c
	y := <- c
	// y, ok := <- c

	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fib_sample() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func gen_fibonacci(c, q chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- q:
			fmt.Println("quit")
			return
		default:
			// run if no other case is ready
		}
	}
}

func gen_sample() {
	c, q := make(chan int), make(chan int)
	
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		q <- 0
	}()

	gen_fibonacci(c, q)
}

func main() {
	//	say_sample()
	// sum_sample()
	// fib_sample()
	// gen_sample()
	mux_sample()
}

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c * SafeCounter) inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c * SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func mux_sample() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.inc("onekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("onekey"))
}
