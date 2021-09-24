package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	s := `[{"FirstName":"James", "LastName":"Bond", "Age":37},{"FirstName":"Miss", "LastName":"MoneyPenny", "Age":27}]`
	bs := []byte(s)
	fmt.Println(s)
	fmt.Println(string(bs))
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n All people data", people)
	for i, v := range people {
		fmt.Println("\nPerson Number: ", i)
		fmt.Println(v.FirstName, v.LastName, v.Age)
	}
}
