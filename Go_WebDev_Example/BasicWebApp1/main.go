package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err:= fmt.Fprintf(w, "Hello WOrld")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Num of byte written: %d", n))
	})
	_= http.ListenAndServe(":8080",nil)
}
