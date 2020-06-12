package main

import (
	"fmt"
	d "github.com/lwheng/learngo/dog" // This is how you use alias: <alias> <package>
	du "github.com/lwheng/learngo/dog/utils"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  d.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(d.YearsTwo(20))

	d.Bark()
	// d.Woof() // This doesn't work because package 'dog' does not have function Woof
	// du.Bark() // This doesn't work because function Bark is defined in its parent package 'dog'
	du.Woof()
}
