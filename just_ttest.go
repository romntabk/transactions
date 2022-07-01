package main

import (
	"fmt"
)
func main(){
	a = 5
	fmt.Println(a)
}

func init(){
	var a int = 3	
	fmt.Println(a)
}
