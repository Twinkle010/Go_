package main

import (
	"fmt"
	"os"
	"io"
)


func main() {
	ReadFile(os.Args[1])
}

func ReadFile(f string) {
	output, err:= os.Open(f)
	fmt.Println(err)
	io.Copy(os.Stdout, output)
}