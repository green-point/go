package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	// fmt.Println(stringst.Kmf("ababababca"))
	// fmt.Println("Hello World")
	// basic_sample()
	defer_sample()
	// struct_sample()
	// array_sample()
	// slice_sample()
	// make2d_array(2, 3)
	// map_sample()
	// funcval_sample()
	// closure_sample()
	// fmt.Printf(stringutil.Reverse("!oG, olleH"))
}

func myPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}

func pointer_sample() {
	var p *int
	i := 42
	p = &i

	*p = 21
	fmt.Println(i)
}

// multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

func basic_sample() {
	sum := 0
	for sum < 10 { // for as a while statement
		sum++
	}

	for {
		if sum > 12 {
			break
		} else {
			sum++
		}

		if v := sum - 1; v > 11 { // if with a init statement
			break
		}
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	switch { // just as switch true
	case sum == 0:
	case sum < 10:
	case sum == 15:
	default:
	}
}

func defer_sample() {
	fmt.Println("counting")

	// defer will be executed after function returned in reverse order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func struct_sample() {
	v := Vertex{2, 4}
	v.X = 4
	fmt.Println(v.X)

	// p := &v
	p.X = 1e9 // or (*p).X
	fmt.Println(*p)
}

func array_sample() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice
	var s []int = primes[1:4]
	fmt.Println(s, len(s))

	// slice share data with the underlying array
	s[0] = 17
	fmt.Println(primes)
}

func slice_sample() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	// r is a slice, reference to the underlying [6]bool array
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// s is a slice, reference to the [6]struct array
	q := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(q)

	// index
	fmt.Println(s[0:6], s[0:], s[:6], s[:])

	// cap and len
	s = s[:0] // slice the slice to give it zero length
	print_slice(s)

	s = s[:4] // extend its length
	print_slice(s)

	s = s[2:] // drop its first two values
	print_slice(s)

	s = s[:4] // before slice, it start from 2, max cap is 4
	print_slice(s)

	var np []int
	fmt.Println(np == nil)

	// make slice
	a := make([]int, 5) // len(a)=5
	print_slice(a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	print_slice(b)

	// b[1] = 1 // panic, since len(b) is 0

	c := b[:2] // len(c)=2, cap(c)=5, extended
	print_slice(c)

	d := c[2:5] // len(d)=3, cap(d)=3
	print_slice(d)

	// slice of slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append element to slice
	var p []int
	print_slice(p)

	p = append(p, 0) // append value on nil slices
	print_slice(p)

	p = append(p, 1) // slice grows as needed
	print_slice(p)

	p = append(p, 2, 3, 4) // append many values at a time
	print_slice(p)
}

func print_slice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func range_sample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}

func make2d_array(dx, dy int) {
	a := make([][]uint8, dy)

	for i := 0; i < len(a); i++ {
		a[i] = make([]uint8, dx)
	}

	fmt.Printf("dx=%d, dy=%d, len(a)=%d, len(a[0])=%d", dx, dy, len(a), len(a[0]))
}

func map_sample() {

	m := map[string]Vertex{
		"Goole": Vertex{
			0, 1,
		},

		"Microsoft": {2, 3}, // ignore the Vertex type declare
	}

	m["Bell Labs"] = Vertex{1, 2}

	fmt.Println(m["Bell Labs"])

	// var n map[string]int
	// n["str"] = 2 // ERROR: nil map

	n := make(map[string]int)
	n["a"] = 1
	n["b"] = 2
	delete(n, "a")

	k := n["b"]
	fmt.Println("n[\"b\"]=", k, "n[\"a\"]=", n["a"])
	v, ok := m["a"]
	fmt.Println("The value:", v, "exist?", ok)
}

func printf_sample() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}

func word_count(s string) map[string]int {
	m := make(map[string]int)

	for _, w := range strings.Fields(s) {
		m[w] += 1
	}

	return m
}

func compute(calculator func(float64, float64) float64) float64 {
	return calculator(3, 4)
}

func funcval_sample() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		ret := f1
		f1 = f2
		f2 = f1 + ret

		return ret
	}
}

func closure_sample() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
