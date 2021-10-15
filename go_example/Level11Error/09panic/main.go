package main

import (
	"os"
)

func main() {
	f, err := os.Open("no-file.txt")
	if err != nil {
		//	fmt.Println("error happened ", err)
		//	log.Println("error happened ", err)
		panic(err)
	}
	f.Close()
}
//panicln is equivalent to Println() followed by call to panic()
// fatalln is equivalent to Println() followed by call to os.exit(1)
//panic function call panic after writing log message
