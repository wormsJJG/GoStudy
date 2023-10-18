package main

import (
	"fmt"
	"strconv"
)

func main() {

	var name string = "정재근"
	convertToIntroduce(&name)
	fmt.Println(name)

	a := "5"

	fmt.Println(strconv.ParseInt(a, 10, 32))
	fmt.Printf("%T\n", a)
}

func convertToIntroduce(nameAddress *string) {

	*nameAddress = "안녕하세요 저는 정재근입니다."
}
