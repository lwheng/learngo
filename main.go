package main

import (
  "github.com/lwheng/learngo/dog"
  "fmt"
)

func main() {
	age := 21
	fmt.Println("John is", age, "years old. In dog years he would be", dog.Years(age), "years old.")
}
