package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	// rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1

	fmt.Println("Guess number (1 ~ 100) : ")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {

		fmt.Println("U chace :", 10-i)
		fmt.Println("Guess number : ")
		inputNumberString, err := reader.ReadString('\n')

		if err != nil {

			log.Fatal(err)
		}

		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)

		if err != nil {

			log.Fatal(err)
		}

		if inputNumber < answer {

			fmt.Println("Ur guess number is lower than answer")
		} else if inputNumber > answer {

			fmt.Println("Ur guess number is heigher than answer")
		} else {

			fmt.Println("Gread U got the number Congratulations")
			break
		}
	}
}
