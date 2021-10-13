package main

import "fmt"
func main(){
	speed,force:=100,2.5
	speed=speed*int(force)
	//this will not give correct result
	fmt.Println(speed)
	speed= int(float64(speed)*force)
	//this will give expected result
	fmt.Println(speed)
}
