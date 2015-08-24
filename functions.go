package main 

import "fmt"

// Takes two ints and returns their sum as an int.Requires explicit returns
func plus(a int, b int )int {

	return a + b
}

func plusPlus(a, b, c int)int {
	
	return a + b + c
}
func main() {

	result := plus(1 , 2)
	fmt.Println("result = ", result)
	
	result = plusPlus(1, 2 ,3)
	fmt.Println("result = ", result);

	/* 
	* Compiler Error
	* too many arguments in call to plusPlus
	*/ 
	
	//result = plusPlus(1, 2 ,3, 4)
}