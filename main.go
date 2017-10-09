package main

import "fmt"

type myStruct struct {
	myVar int
	myOtherVar string
}

func main() {
	x:=100
	x*=2
	y :=x/ 3
	z:=50

	fmt.Print(y)
}
