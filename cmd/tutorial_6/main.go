package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Printf("can make it")
	} else {
		fmt.Printf("can not make it")
	}
}
func main() {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"Alex"}} //this is sigining the value
	// myEngine.mpg = 24                                                                  //this too
	// myEngine.gallons = 25                                                              //this too
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	// // fmt.Printf("Total miles left in tank: %V", myEngine.milesLeft())
	var myEngine electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine, 50)
}
