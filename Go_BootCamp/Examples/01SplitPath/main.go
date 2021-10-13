package main

import (
	"fmt"
	"path"
)

func main() {
	dir,file := path.Split("Golang/Go/Src/GoHelp.txt")
	fmt.Println("Dir:",dir)
	fmt.Println("File:",file)
}
