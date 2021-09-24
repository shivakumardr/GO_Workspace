package main

import (
	"fmt"
	"math"
)

/*
Hands-on exercise #5
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
code: https://play.golang.org/p/NGGikHNvHv

*/
type circle struct {
	radius float64
}
type square struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	crl := circle{
		radius: 6.6,
	}
	sqr := square{
		length: 4.4,
		width:  4.4,
	}
	info(crl)
	info(sqr)
}
