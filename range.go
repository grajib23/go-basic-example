package main 

import "fmt"

func main() {

	//declear and initialize nums array 
	nums := []int{1, 2, 3}
	sum := 0

	// _ optional parameter.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	for i, num := range nums {
		if i ==2 {
			fmt.Printf("i = %d and num = %d\n", i, num)
		}
	}

	m := map[string]int{"key1" : 1, "key2" : 2, "key3" : 3}
	for key, value := range m {
		fmt.Printf("key = %s and value = %d\n", key, value)
	}

	// range on strings iterates over Unicode code points.
	for i,v := range "AB" {
		fmt.Printf("i = %d and v = %d\n", i, v)
	}
}