package main 

import "fmt"


// variadic 
func sum(nums... int)int {

	result := 0
	for i,num := range nums {
		fmt.Println("index = ", i)
		result += num
	}
	return result
}

func main() {
	
	result := sum(1, 2)
	fmt.Println("Result is = ", result)
	result = sum(1, 2, 3, 4)
	fmt.Println("Result is = ", result)
}