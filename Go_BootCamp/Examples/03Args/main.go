package main

import (
	"fmt"
	"os"
)
func main(){
	fmt.Printf("\n Command line args %#v\n", os.Args)
	fmt.Println("Path:",os.Args[0])
	fmt.Println("First argument:",os.Args[1])
	fmt.Println("Second argument:",os.Args[2])
	fmt.Println("Third argument:",os.Args[3])
}
