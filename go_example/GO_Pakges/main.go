package main

import (
	"github.comGo_workspacego_exampleGO_Pakges/helpers"
	"log"
)

const numPool = 1000
func CalculateValue(intChan chan int){
	randomNum:= helpers.RandomeNum(numPool)
	intChan <- randomNum
}
func main() {
intChan := make(chan int)
defer close(intChan)
go CalculateValue(intChan)
num:= <- intChan
log.Println(num)
}
