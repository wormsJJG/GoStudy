package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0}

	s[4] = 99
	s[2] = 11

	for i := 0; i < len(s); i++ {

		fmt.Println(s[i])
	}

	copyS := s[1:4]

	for _, value := range copyS {

		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "students"}
	testS := test[:2]
	testS2 := test[1:]
	fmt.Println(testS2[1])
	fmt.Println(len(testS))

	a := make([]string, 4, 5)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"

	// as := a[0:2]
	c := append(a, "y", "x")

	fmt.Println(a, len(a), cap(a))
	fmt.Printf("%x\n", &a)
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x\n", &c)
}
