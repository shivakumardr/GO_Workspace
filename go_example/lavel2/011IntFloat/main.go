package main

import "fmt"

var a uint8
var b uint16
var c uint32
var d uint64
var e int8
var f int16
var g int32
var h int64
var i float32
var j float64

func main() {
	a = 2
	b = 300
	c = 3000
	d = 40000
	e = -100
	f = -3000
	g = -10000
	h = -87132847912833
	i = -300.98980
	j = 49909.37432423442

	k := 234234
	l := 23423.877898987
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println(g)
	fmt.Printf("%T\n", g)
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println(j)
	fmt.Printf("%T\n", j)
	fmt.Println(k)
	fmt.Printf("%T\n", k)
	fmt.Println(l)
	fmt.Printf("%T\n", l)

}
