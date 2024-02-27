package main

import "fmt" 

func main() {
	fmt.Printf("%s\n", Hi("world"))
}

func Hi(name string) string {
   return fmt.Sprintf("Hi, %s", name)
}