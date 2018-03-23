package main

import (
	"strings"
	"time"
	"math"
	"fmt"
	"io"
)

func main() {
	// method_sample() // method call
	// interface_sample() // interface sample
	reader_sample()
}

// methods
type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type PVertex *Vertex

// pointer receiver will change the input struct
// same as func scale(v *Vertex, f float64)
func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func method_sample() {
	
	// method of vertex
	v := Vertex {X:1, Y:2}
	fmt.Println(v.abs())

	// pointer receiver
	v.scale(2.0)

	p := &v
	p.scale(3.0) // same as v.sacle(3.0)
	scale(p, 2.0)
	
	q := &Vertex {3, 4}
	q.abs()
} 

type Abser interface {
	abs() float64
}

type ufloat float64

func (f *ufloat) abs() float64 {
	return math.Abs(float64(*f))
}

type I interface {
	M()
}

type T struct {
	S string
}

// func (t T) M() {
//	fmt.Println(t.S)
//}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func any_describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interface_sample() {
	var a Abser
	v := Vertex{3, 4}
	a = v 
	fmt.Println(a.abs())
	
	f := ufloat(2.0)
	a = &f // only *ufloat64 implemented abs interface 
	fmt.Println(a.abs())

	// implicitly implment
	var i I = &T{"hello"}
	i.M()
	describe(i) // inteface values is (value, type)

	// nil value interface
	var n I
	var t *T
	n = t
	describe(n)
	n.M() // interface has been asigned to T, so it is not nil, just the t value is nil 

	var k interface{}
	any_describe(k)

	k = 42
	any_describe(k)

	k = "Hello"
	any_describe(k)

	// type assertions
	s := k.(string) // if k is not a string, program panic
	fmt.Println(s)

	s, ok := k.(string) // if k is not a string, ok is false, s is zero value
	fmt.Println(s, ok)

	// type switch
	action_by_type(21)
	action_by_type("hello")
	action_by_type(true)

	// struct stringers
	ll := Person{"Leo", 41}
	fmt.Println(ll)
}

func action_by_type(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v + 1)
	case string:
		fmt.Println(len(v))
	default:
		fmt.Printf("<UNT>: %T\n", v)
	}
}

type Person struct {
	Name string
	Age int
}

// Stringers
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
type error interface {
	Error() string
}
*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError { time.Now(), "panic"};
}

func error_sample() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func reader_sample() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}