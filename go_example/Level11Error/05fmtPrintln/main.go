package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("no-file.txt")
	if err != nil {
	//	fmt.Println("error happened ", err)
		log.Println("error happened ", err)

	}
	f.Close()
}
