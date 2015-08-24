package main 

import "fmt"

func main() {
	
	m := make(map[string]int)
	fmt.Println("m = ", m)
	fmt.Println("len(m)", len(m))

	m["key1"] = 7
	m["key2"] = 8
	fmt.Println("m = ", m["key"])
	fmt.Println("len(m)", len(m))
	fmt.Println("m[key1] = ", m["key1"])
	
	/**
	 * throw Exception
	 * cannot use "test" (type string) as type int in assignment
	*/
	// m["key3"] = "test"


	 /**
	 * throw Exception
	 * cannot use 7 (type int) as type string in assignment
	*/

	 // m1 := make(map[string]string)
	 // m1["key1"] = 7

	 // delete key 
	delete(m, "key1")
	fmt.Println("m = ", m)

	// The optional second return value
	_,pros1 := m["key1"]
	fmt.Println("pros1 = ", pros1)
	_,pros2 := m["key2"]
	fmt.Println("pro2 = ", pros2)

	// declare and initialize a new map in the same line 
	m2 := map[string]int {"key1" : 1, "key2": 2, "key3" : 3}
	fmt.Println("m2 = ", m2)
}