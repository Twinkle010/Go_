package main

import "fmt"

// struct is a ds in go that stores different data typess
// struct is a kind of dict in python

///embedding structs
type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo // or give just contactInfo that go will consider as a field contactInfo of type contactInfo
}
// values assigned when var is just initialised and not assigned
// Type ZeroValue
// string ""
// int 0
// float 0
// bool False

func main() {
	// alex := person{"Alex", "Ander"} //assumes in order
	// or use as below
	// alex := person{firstName: "Alex", lastName: "Ander"}
	// fmt.Println(alex)
	// or
	var alex person
	alex.firstName ="Alex"
	fmt.Println(alex)
	fmt.Printf("%+v", alex) //print all the fields in alex
	twinkle := person{
		firstName: "Twinkle",
		lastName: "Set",
		contact: contactInfo{
			email: "twinkle@gmail.com",
			zipCode: 123456, //when defining multiple fields instructs, every single line must contain a comma even if it's last line
		},
	}
	// fmt.Println(twinkle)
	// fmt.Printf("%+v", twinkle)
	fmt.Println("---")
	twinkle.updatefirstName1("Lilly") // but frst name is not updated
	// cause go is a pass by value lang, the value is copied and stored as some new struct at a different physical location when called as a receiver i.e whenver a value is passed to a fn and the new struct is updated when update fn is called, so the original remains same 
	//applicable for all data types
	// then what to do??
	// use pointers

	twinkle.print()
	fmt.Println("---")
	twinklePointer := &twinkle
	twinklePointer.updatefirstName("Lilly") //create a pointer and use pointer
	twinkle.print()
	fmt.Println("---")
	twinkle.updatefirstName1("Lilly")
	twinkle.print()
	fmt.Println("---")
	fmt.Println("references")
	mySlice := []string{"Hi", "There", "!"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice) //hi is updated to bye here without using pointers anywhere, whyy??
	// when a slice is intialised, internally a structure is created with params lenght, capacity, ptr to head where in internally refers to an array
	// so when a slice is passed as an arg to a fn , it creates a copy in memory like other ds, but the ptr to head still refers to the origin value, hence it's updated
	//ValueTypes ReferenceTypes
	// int, float, bool, string, structs -> use pointers to change things in fn
	//slices, maps, channels, pointers, functions -> no need of pointers
}


// struct receiver fns
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updatefirstName1(newName string) {
	p.firstName = newName
}

// &var: give me the memory address of the value this var is pointing at
// *pointer: give me the value this memory address is pointing at
// turn address into value with *address
// turn value into address with &value
//*pointershortcut* instead of getting address and doing things always 
// apart from the above method, if the receiver fn is of type pointer to person, you can call the fn with type person, go will automatically turn person to pointertoPerson
func (pointerToPerson *person) updatefirstName(newName string) { // *person here's a type description i.e we're working with a pointer to a person
	(*pointerToPerson).firstName = newName
}


func updateSlice(s []string){
	s[0] = "Bye"
}