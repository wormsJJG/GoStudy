package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, err := reader.ReadString('\n')
	if err != nil {

		log.Fatal(err)
	}

	inputScore = strings.TrimSpace(inputScore)
	score, err := strconv.ParseFloat(inputScore, 64)
	var grade string

	if score >= 90 {

		grade = "A grade"
	} else {

		grade = "under A gorade"
	}

	fmt.Println(grade)
	fmt.Println("You will get", grade)
}
