package main 

import "fmt"
import "math"

const s string = "test string"

func main() {
	fmt.Println("s = ", s);
	const n = 500000000
	fmt.Println("n = ", n)
	const d = 3e20 / n
	fmt.Println("d = ", d)
	fmt.Println("int64(d) = ", int64(d))
	fmt.Println("math.Sin(n) = ", math.Sin(n))	
}