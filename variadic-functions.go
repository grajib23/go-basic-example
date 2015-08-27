package main 

import "fmt"


// variadic 
func sum(nums... int)int {

	result := 0
	for _,num := range nums {
		result += num
	}
	return result
}

func main() {
	
	result := sum(1, 2)
	fmt.Println("Result is = ", result)
	result = sum(1, 2, 3, 4)
	fmt.Println("Result is = ", result)
	nums := []int{1, 2, 3, 4, 5}
	result = sum(nums...)
	fmt.Println("Result is = ", result)
}