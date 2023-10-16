package main

import (
	"fmt"
	"strings"
)

func main() {

	brokenWords := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")
	fixedWrods := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWrods)
} 