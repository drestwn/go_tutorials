package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello world"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision2(numerator, denominator)

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {

		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	// this is new.
	// if condition 1==2 && 2==2, after checking 1==2 and its true, 2==2 is skipped, Then whats the point of && if the second is not checked first? still no idea

	// switch cases of the ifs...

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)

	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)

	}

	//you can use switch with checcking the value, this is new
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")

	default:
		fmt.Printf("The division was not close")

	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// if wanted a function to return a value, need to specify what type of the return (1 is like this)
func intDivision(numerator int, deminator int) int {
	var result int = numerator / deminator
	return result
}

// if wanted a function to return a value, need to specify what type of the return (2 is like this)
func intDivision2(numerator int, denominator int) (int, int, error) {

	// can create type error value the default is nil
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
