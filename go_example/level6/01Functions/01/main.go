package main

import "fmt"

func main() {
	fmt.Println("Good Morning All")
	foo()
	woo("Where are you")
	s1 := bar("Hello")
	fmt.Println(s1)
	x, y := funcreturntwo("sending")
	fmt.Println(x, y)
	mp := returningmap()
	fmt.Println(mp)
	for i, v := range mp {
		fmt.Println(i, v)
	}
	strs, slc := returningSlice("Shiva Says")
	fmt.Println(strs)
	fmt.Println(slc)
	for i, v := range slc {
		fmt.Println(i, v)
	}

}
func foo() {
	fmt.Println("I am in foo")
}
func woo(s string) {
	fmt.Println("I am in woo, says", s)
}
func bar(s string) string {
	return fmt.Sprint("I am in Bar, syas", s)
}
func funcreturntwo(st string) (string, bool) {
	a := fmt.Sprint("Returning two val", st)
	b := true
	return a, b
}
func returningmap() map[string]int {
	m := map[string]int{
		"Shiv": 999,
		"Nis":  998,
		"Push": 997,
	}
	return m
}
func returningSlice(s string) (string, []string) {
	sl := []string{
		"Bang",
		"Dvg",
		"Shyg",
		"Kvrt",
		"Bhp",

	}
	return s, sl
}
