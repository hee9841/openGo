package main

import "fmt"

type part struct {
	description string
	count int
}
func showInfo(p part){
	fmt.Println("Description : ", p.description)
	fmt.Println("count : " ,p.count)
}
func main(){
	var bolts part
	bolts.description = "hex bolts"
	bolts.count = 100
	shwoInfo(bolts)
}
