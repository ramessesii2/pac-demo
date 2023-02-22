package main

import "fmt"

const unused = `unusedstr`

func main() {
	fmt.Println("Hello world")
	fmt.Printf("Hello world using %s", unused)
}
