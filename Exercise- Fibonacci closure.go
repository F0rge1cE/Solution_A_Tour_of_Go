package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := -1
	current := 1
	next := 0
	
	return func () int{
		
		next = prev + current
		prev = current
		current = next
		
		
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
