package main 

import "fmt"

/*
* anonymous functions, which can form closures.Anonymous functions are
* useful when you want to define a function inline without having to name it.
*/
func intSeq() func()int {
	i := 0
	return func()int {
		i++
		return i
	}
}
func main() {
	intNext := intSeq();
	fmt.Println(intNext());
	fmt.Println(intNext());
	fmt.Println(intNext());
	
	intOther := intSeq();
	fmt.Println(intOther());
}