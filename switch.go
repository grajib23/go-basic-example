package main 

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write ", i , "as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("this is weekend")
	default:
		fmt.Println("this is weekday");
	}

	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("this is before 12")
	default:
		fmt.Println("this is after 12")
	}
}