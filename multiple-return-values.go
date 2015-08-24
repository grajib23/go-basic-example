package main 

import "fmt"

/*
* Go has built-in support for multiple return values.
* The (int, int) in this function signature shows that
* the function returns 2 ints.
*/ 
func addSub(a , b int) (int, int) {

	return (a + b), (a-b)
}

func main() {

	// multiple assignment.
	add,sub := addSub(10, 5)
	fmt.Printf("add = %d , sub = %d\n", add, sub)

	// Use the blank identifier _ for subset of the returned values
	_, sub = addSub(11, 1)
	fmt.Println("sub = ", sub)
}