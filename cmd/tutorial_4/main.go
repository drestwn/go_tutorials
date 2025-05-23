package main

import (
	"fmt"
)

func main() {

	//arrays
	var intArr [3]int32
	fmt.Println(intArr[0])   //0
	fmt.Println(intArr[0:3]) //[0, 0]

	intArr[1] = 123
	fmt.Println(intArr[0:3]) //123, 0

	//can type like this
	var intArr2 [3]int32 = [3]int32{1, 2, 3} //3 is the capacity
	fmt.Println(intArr2)

	//slices
	var intSlices []int32 = []int32{4, 5, 6}
	fmt.Println(intSlices)
	intSlices = append(intSlices, 7) //inserting 7 to the array

	var intSlices2 []int32 = []int32{8, 9}
	intSlices = append(intSlices, 7)             //inserting 7 to the array
	intSlices = append(intSlices, intSlices2...) //inserting 7 to the array

	fmt.Println(intSlices) //4,5,6,7

	//anotherway define to make array
	var intSlice3 []int32 = make([]int32, 3, 8) //3 is length, 8 is capacity. length only will make default 3
	fmt.Println(intSlice3)

	//maps

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 12, "Sarah": 45}

	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // uint8 return 0 as default. Map always return a value so it will be 0 even its not showing
	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("the age of Adam is %v\n", age)
	} else {
		fmt.Println("name is not existed")
	}

	// delete(myMap2, "Adam") //to delete

	//loops !!

	for name := range myMap2 {
		fmt.Printf("Name %v\n", name)
	}
	for name, age := range myMap2 {
		fmt.Printf("Name %v, age:%v \n", name, age)
	}

	//for array
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	//no "while", use this

	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	//or
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	// i-- (decrement)
	// i++ (increment)
	// i+=10 (increment by 10)
	// i-=10 (decrement by 10)
	// i*=10 (times 10)
	// i/=10 (devided 10)

}
