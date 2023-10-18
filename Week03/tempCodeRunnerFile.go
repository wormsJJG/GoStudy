package main

import (
	"fmt" //모듈
	"strconv"
)

func main() {

	a := "5"

	b, _ := strconv.Atoi(a)

	fmt.Printf("%T\n", b)
	println("Hello Go!")                           //function                //줄 바꿔서 출력 Hello Go!
	fmt.Print("OpenSource Programming~\t", "GO\n") // OpenSource Programming~GO
	fmt.Print("asdasdasd")
	fmt.Println("asdasdasdasdasdasd") // fmt 안에 이 함수가 있다. <- method
	fmt.Println("asdasdasdasdasdasd")
}
