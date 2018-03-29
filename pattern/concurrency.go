package main

import (
	"time"
	"fmt"
	"math/rand"
)

// generator

func boring(msg string) <-chan string {
	// return receive-only channel string
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func generator_sample() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

// multiplexing
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()

	return c
}

func multiplex_sample() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring.");
}

// restoring sequencing

type Message struct {
	str string
	wait chan bool
}

func boring_wait(waitForIt chan bool, c chan Message, msg string) {
	for i := 0; ; i++ {
		c <- Message { fmt.Sprintf("%s: %d", msg, i), waitForIt }
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		<- waitForIt
	}
}

func restore_sample() {
	c := make(chan Message)

	waitForIt := make(chan bool)
	
	go boring_wait(waitForIt, c, "Joe")
	go boring_wait(waitForIt, c, "Ann")

	for i := 0; i < 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}

func main() {

}

// select

func select_fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <-input2: c <- s
			}
		}
	}()

	return c
}

func timeout_sample(quit chan string) {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You're too slow")
			return
		case <-quit:
			// cleanup()
			quit <- "See you."
			return
		}
	}
}

// daisy chain

func f(left, right chan int) {
	left <- 1 + <- right
}

// build 1000 channels like: l <- r <- r ... <- r
func daisy_sample() {
	const n = 1000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
