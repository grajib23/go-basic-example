package main 

import "fmt"

func main() {
	var a[5]int
	fmt.Println("a = ", a)

	a[4] = 100
	fmt.Println("a = ", a)
	fmt.Println("a[4] = ", a[4])
	fmt.Println("len(a) = ", len(a))
	b := [5]int{1,2,3,4,5}
	fmt.Println("b = ", b)

	var twoD[3][2]int
	for i := 0; i < 3 ; i++ {
		for j := 0 ; j < 2; j ++ {
			twoD[i][j] = i + j
		}
	} 
	fmt.Println("twoD = ", twoD)
	fmt.Println("len(twoD) = ", len(twoD))
}