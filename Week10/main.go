package main

import (
	"fmt"

	"Week10/src/greeting/greeting"

	"Week10/mymath/mymath"
)

func main() {

	greeting.Hello()
	greeting.Hi()
	fmt.Println(mymath.MyAbs(3))
}
