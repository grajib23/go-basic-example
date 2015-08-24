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
			fmt.Println("i and num = ", i, num)
		}
	}

	m := map[string]int{"key1" : 1, "key2" : 2, "key3" : 3}
	for key, value := range m {
		fmt.Println("key and value pf m = ", key, value)
	}

	for k,v := range "AB" {
		fmt.Println("k and v = ", k, v)
	}
}