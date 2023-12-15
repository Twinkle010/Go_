// wkt every value has a type
// every fn has to specify type of args
//but does that mean, we have to use different fns for diff types even if the logic is identical??
package main

import "fmt"

type bot interface{
	getGreeting() string // if any other type has this fn defined, then it's also a member of this type bot , so these types can now access printGreeting fn
	//rules of interface: mention the fn name, arg types and return types here
}
//anything that matches description of type above is of type bot
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hello World!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

//similar fns below , throws error
//  cannot use sb (variable of type spanishBot) as englishBot value in argument to printGreeting
// ./main.go:32:6: printGreeting redeclared in this block
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}