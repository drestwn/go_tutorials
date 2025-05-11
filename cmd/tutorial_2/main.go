package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32) //this is needed because cannot do + on different type
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello world" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("test")) //to console.log

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3) //there is some default value of each variable not assigned value 0, " ", false

	myVar := "text"
	var myVar2 = "abc"
	fmt.Println(myVar, myVar2) //u dont have to type type of it. there is some default

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// const works too. But you cannot change value or reassign it
	const myConst string = "const value"
	fmt.Println(myConst)

	//example like pi value
	const pi float32 = 3.1415
	fmt.Println(pi)
}
