package main 

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("i = ", i);
		i++
	}

	for j := 1; j <=3 ; j++ {
		fmt.Println("j = ", j)
	}

	for {
		fmt.Println("loop");
		break;
	}
}