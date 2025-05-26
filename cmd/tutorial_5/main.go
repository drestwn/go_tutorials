package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "Résumé" //Go see string as UTF-8 encoding

	var indexed = myString[0]
	fmt.Println(indexed)

	//so if need to see the length like in JS,

	var myString2 = []rune("Résumé")
	// var indexed2 = myString2[1]

	fmt.Printf("\nThe length of my string2 is %v", len(myString2))

	//the variable of string not able to concate, use build in function by go

	var strSlice = []string{"S", "U", "B", "S", "C", "R", "I", "B", "E"}
	var strBuilder strings.Builder //this
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])

	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
