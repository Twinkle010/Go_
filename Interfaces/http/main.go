package main


import (
	"fmt"
	"net/http"
)

func main() {
	resp, errorc := http.Get("http://google.com") //Get fn returns   a pointer of response and error if any
	fmt.Println("Response: ", *resp)
	fmt.Println("Error: ", errorc)
}