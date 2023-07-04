package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty ", s, "len", len(s), "capacity", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set are ", s)
	fmt.Println("getting the element at index ", s[2])
	fmt.Println("length ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copying ", c)

	l := s[0:3]
	fmt.Println("sl1", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declaration ", t)

	twod := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twod[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println(twod)
}
