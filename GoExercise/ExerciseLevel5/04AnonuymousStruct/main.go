package main

import "fmt"

func main() {
	p1 := struct {
		first   string
		last    string
		age     int
		friends map[string]int
		drinks  []string
	}{
		first: "James",
		last:  "Bond",
		age:   34,
		friends: map[string]int{
			"prk": 9999,
			"vrn": 777,
			"snd": 888,
		},
		drinks: []string{
			"Bear",
			"water",
			"coock",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
	for i, v := range p1.friends {
		fmt.Println(i, v)
	}
	for i, v := range p1.drinks {
		fmt.Println(i, v)
	}

}
