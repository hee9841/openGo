package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description : ", p.description)
	fmt.Println("count : ", p.count)
}
func minmumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	p := minmumOrder("hex bolts")
	fmt.Println(p.description, p.count)
}
