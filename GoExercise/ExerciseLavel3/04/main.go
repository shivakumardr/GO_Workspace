package main

import "fmt"
/*
Hands-on exercise #4
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
solution: https://play.golang.org/p/9VpnB-I1Pz
*/
func main(){
	bd := 1983
	for {
		if(bd >2021){
			break
		}
		fmt.Println(bd)
		bd++
	}
}
