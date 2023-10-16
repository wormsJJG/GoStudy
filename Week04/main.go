package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// 기본값 제로벨류
	var a int = 9
	var b = 3.14
	var c, d string
	var e float32 = 3.14
	var f string // 기본 값 빈 문자열
	var g bool   // 기본 값 flase
	// a = 9
	c = "inha~"
	d = "GO..."
	h := 'A' // rune 은 int32
	i := "문자열"
	J := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음"

	fmt.Println(float64(a) > b)
	fmt.Println(float64(a) * b)
	fmt.Println(a, b, c, d, e, f, g, h, i, J)
	fmt.Println(reflect.TypeOf(b)) // 기본 활당 자료형은 float64

	fmt.Println(reflect.TypeOf('B')) //rune 은 int32
	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf(2.71))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf("GO!"))
	fmt.Println('A') //rune (unicode)
	fmt.Println('정', '재', '근')
	fmt.Println(math.Floor(2.75))
	fmt.Println(math.Ceil(2.75))
	fmt.Println(strings.Title("오픈소스\t프로그래밍~\nGo"))
	fmt.Println(strings.Title("open souce programming go!"))
}
