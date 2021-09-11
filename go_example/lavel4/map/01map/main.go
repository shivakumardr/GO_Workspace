package main

import "fmt"

func main() {
	m := map[string]int{
		"Davangere": 577002,
		"Bangalore": 560001,
		"RR Nagar":  560098,
	}
	fmt.Println(m)
	fmt.Println(m["Davangere"])
	fmt.Println(m["Shyagale"])
	v, ok := m["shyagale"]
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = m["RR Nagar"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok = m["shyagale"]; ok {
		fmt.Println("This is printng", m["shyagale"])
	}
	if v, ok = m["Bangalore"]; ok {
		fmt.Println("This is printng", m["Bangalore"])
	}
	m["shyagale"] = 57702
	if v, ok = m["shyagale"]; ok {
		fmt.Println("This is printng", m["shyagale"])
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// slice
	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
