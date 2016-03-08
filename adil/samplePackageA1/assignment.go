package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func mul(x int, y int) int {
	return x * y
}

func main() {
  fmt.Println("Add 2 2  = : ")
	fmt.Println(add(2, 2))
    fmt.Println("Add 4 6  = : ")
  	fmt.Println(add(4, 6))
      fmt.Println("Mul 2 2  = : ")
    	fmt.Println(mul(2, 2))
}
