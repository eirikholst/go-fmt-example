package main

import "fmt"

type myStruct struct {
	myVar int
	myOtherVar string
}

func main(){
	x	:=	100
	x	*=2
	y:=x/3

	a := myStruct {
		myVar:2,
		myOtherVar:"two",
	}
	b := a.myVar
	c := b+2

	fmt.Println(y * c)
}
