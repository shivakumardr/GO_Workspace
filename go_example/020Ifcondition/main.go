package main

import "fmt"
func main(){
	if true{
		fmt.Println("condition is true")
	}
	if false{
		fmt.Println("condition is false")
	}
	if !true{
		fmt.Println("it's not of true (means false)")
	}
	if !false{
		fmt.Println("it's not of false (means true)")
	}
	if x:= 10; x>10{
		fmt.Println("This number is more than 10")
	}else if  x < 10{
		fmt.Println("This number is less than 10")
	}else{
		fmt.Println("The number is 10")
	}
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

}
