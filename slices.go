package main 

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emt s =  ", s)
	fmt.Println("len(s) = ", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set s = ", s)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append s = ", s)
	fmt.Println("append s length = ", len(s))

	c := make([]string, len(s))
	fmt.Println("emt c=  ", c)
	fmt.Println("len(c) = ", len(c))

	copy(c , s)
	fmt.Println("c =  ", c)

	l := s[2:5]
	fmt.Println("sl1 = ", l)
	fmt.Println("len(l) = ", len(l))

	l = s[2:]
	fmt.Println("sl2 = ", l)
	fmt.Println("len(l) = ", len(l))

	l = s[:5]
	fmt.Println("sl3 = ", l)
	fmt.Println("len(l) = ", len(l))

	t := []string{"g", "h", "i"}
	fmt.Println("t = ", t)
	fmt.Println("len(t) = ", len(t))

	k := t[1:]
	fmt.Println("k = ", k)
	fmt.Println("len(k) = ", len(k))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1;
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}