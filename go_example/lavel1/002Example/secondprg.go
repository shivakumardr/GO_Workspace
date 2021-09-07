package main

import "fmt"

func main() {
	fmt.Println("This is secon progam")
	foo()
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
		fmt.Println("\t", i)
	}

	bar()
}
func foo() {
	fmt.Println("i am in foo")
}
func bar() {
	fmt.Println("i am in bar exiting program")

}
