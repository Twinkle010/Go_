package main

import "fmt"

type number []int 

func main() {
	output := newList()
	fmt.Println(output)
	fmt.Println("-----")
	for value := range(output){
		checkType(value)
	}
	// checkType(output)
}


func newList() number{
	newSlice := []int{0,1,2,3,4,5,6,7,8,9,10}
	// for value := range [10] int{
	// 	newSlice =  append(newSlice, value)
	// }
	return newSlice
}

func checkType(v int) {
	// for v := range [i] []int{
	if(v%2==0){
		fmt.Println(v,"is even")
	}else{
		fmt.Println(v,"is odd")
	}
}