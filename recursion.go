package main 

import "fmt"

// recursive functions.
func fact(num int)int {
	if num == 0 {
		return 1
	}
	return num * fact(num - 1)
}

func main() {
	result := fact(3)
	fmt.Println("Result is = ", result)
}