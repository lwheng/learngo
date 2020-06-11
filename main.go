package main

import (
	"fmt"
	d "github.com/lwheng/learngo/dog" // This is how you use alias: <alias> <package>
	du "github.com/lwheng/learngo/dog/utils"
)

func main() {
	age := 21
	fmt.Println("John is", age, "years old. In dog years he would be", d.Years(age), "years old.")

	d.Bark()
	// d.Woof() // This doesn't work because package 'dog' does not have function Woof
	// du.Bark() // This doesn't work because function Bark is defined in its parent package 'dog'
	du.Woof()
}
