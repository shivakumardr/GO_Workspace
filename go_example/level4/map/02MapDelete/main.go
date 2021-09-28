package main

import "fmt"

func main() {
	m := map[string]int{
		"Bangalore":  560001,
		"Davanagere": 577001,
		"RR Nagara":  560098,
	}
	fmt.Println(m)
	delete(m, "Davanagere") // Deleting from map
	fmt.Println(m)
	delete(m, "shyagale") // deleting from map but this is not in map still not giving error
	fmt.Println(m)
	m["shyagale"] = 577002
	fmt.Println(m)
	fmt.Println(m["Bangalore"])
	fmt.Println(m["shyagale"])
	if v, ok := m["Bangalore"]; ok {
		fmt.Println(v, ok)
		delete(m, "Bangalore")
	}
	fmt.Println(m)

}
