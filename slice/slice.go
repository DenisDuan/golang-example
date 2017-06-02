package main

import "fmt"

func main() {
	// slice initialization in single line
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// Slice initialization
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("Get:", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy of S: ", c)

	// slicing operation
	l := s[2:5]
	fmt.Println("sl[2:5]: ", l)

	l = s[:5]
	fmt.Println("s[:5]: ", l)

	l = s[2:]
	fmt.Println("s[2:]: ", l)

	// the length of slice can be dynamic (vary)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two demensional slice: ", twoD)

}