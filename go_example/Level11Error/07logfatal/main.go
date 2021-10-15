package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	f, err := os.Open("no-file.txt")
	if err != nil {
		//	fmt.Println("error happened ", err)
		//	log.Println("error happened ", err)
		log.Fatalln(err)

	}
	f.Close()
}
func foo() {
	fmt.Println("Checking when log.fatalln runs dfer func")
	fmt.Println("When os.Exit() call,defer func does't call")
}
