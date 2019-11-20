package main

import (
    "fmt"
	"strconv"
	"assignment/loop"
	"assignment/recursion"
	"assignment/reverse"
	

)


func main() {
    for i := 0; i <= 9; i++ {
        fmt.Print(strconv.Itoa(loop.FibonacciUsingLoops(i)) + " ")
    }
    fmt.Println("")
    for i := 0; i <= 9; i++ {
        fmt.Print(strconv.Itoa(recursion.FibonacciUsingRecursion(i)) + " ")
    }
	fmt.Println("")
	
	for i := 0; i <= 9; i++ {
	fibStr :=strconv.Itoa(loop.FibonacciUsingLoops(i))
	}
    strRune := reverse.Runes(fibStr)
    reversedfibStr := strRune.ReverseString()    
    fmt.Println("Reversed: ",reversedfibStr)
}