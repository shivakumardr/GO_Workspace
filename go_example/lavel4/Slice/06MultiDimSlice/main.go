package main

import "fmt"

func main() {
	jb := []string{"shiva", "pushpa", "Pavaki", "parvathi"}
	fmt.Println(jb)
	fmt.Println(len(jb))
	fmt.Println(cap(jb))
	Jd := []string{"Shyagale", "bengaluru", "Davangere", "Karnataka"}
	fmt.Println(Jd)
	fmt.Println(len(Jd))
	fmt.Println(cap(Jd))
	sp := [][]string{jb, Jd}
	fmt.Println(sp)
	fmt.Println(len(sp))
	fmt.Println(cap(sp))
	i := []int{1, 2, 3, 4}
	fmt.Println(i)
	j := []int{5, 6, 7}
	fmt.Println(j)
	k := [][]int{i, j}
	fmt.Println(k)
	fmt.Println(len(k))
	fmt.Println(cap(k))

}
