package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Good evening")
	fmt.Fprintln(os.Stdout, "Hello standard lib")
	io.WriteString(os.Stdout, "Hello standard lib")
}
