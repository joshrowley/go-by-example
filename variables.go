package main

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println("zero values by type:")
	var g int
	fmt.Println("int: ", g)

	var h int64
	fmt.Println("int64: ", h)

	var i float64
	fmt.Println("float64: ", i)

	var j bool
	fmt.Println("bool: ", j)

	var k string
	fmt.Println("string: ", k)

	var l *int
	fmt.Println("pointer: ", l)

	var m User
	fmt.Println("user: ", m)

	// slices
	var n []int
	fmt.Println("int slice: ", n)

	var o []string
	fmt.Println("string slice: ", o)

	// maps
	var p map[string]int
	fmt.Println("string => int map: ", p)

	// channels
	var q chan int
	fmt.Println("int channel: ", q)

	// functions
	var r func()
	fmt.Println("function: ", r)

	// interfaces
	var s interface{}
	fmt.Println("interface: ", s)

	// arrays
	var t [3]int
	fmt.Println("array: ", t)

	// structs
	var u User
	fmt.Println("structs: ", u)

	// time
	var v time.Time
	fmt.Println("time: ", v)
}
