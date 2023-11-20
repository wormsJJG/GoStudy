package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([4]float64, error) {
	var numbers [4]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

// func main() {

// 	// var primes [3]int

// 	// primes[0] = 2
// 	// primes[2] = 5

// 	// for i := 0; i < 3; i++ {

// 	// 	fmt.Println(primes)
// 	// }

// 	// primes := [3]int{2, 3, 5}

// 	// primes[2] = 6

// 	// for i := 0; i < 3; i++ {

// 	// 	fmt.Println(primes)
// 	// }

// 	// test := [5]bool{true, true, true}

// 	// // fmt.Println(test[3])

// 	// i := 0

// 	// for i < len(test) {

// 	// 	fmt.Println(test[i])
// 	// 	i++
// 	// }

// 	// for idx, prime := range primes {

// 	// 	fmt.Println(idx, prime)
// 	// }

// 	file, err := os.Open("data.txt")

// 	if err != nil {

// 		log.Fatal(err)
// 	}

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {

// 		fmt.Println(scanner.Text())
// 	}

// 	err = file.Close()

// 	if err != nil {

// 		log.Fatal(err)
// 	}

// }
