package main

import "fmt"

func main() {

	var name string = "정재근"
	convertToIntroduce(&name)
	fmt.Println(name)
}

func convertToIntroduce(nameAddress *string) {

	*nameAddress = "안녕하세요 저는 정재근입니다."
}
