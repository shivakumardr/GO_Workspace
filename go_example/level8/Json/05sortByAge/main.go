package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func main() {
	p1 := Person{Name: "Parvathi", Age: 6}
	p2 := Person{"Pavaki", 5}
	p3 := Person{"Pushpavathi", 34}
	p4 := Person{"Shivakumar", 33}
	p5 := Person{"ASu", 10}
	people := []Person{p1, p2, p3, p4, p5}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("####************************####")
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
