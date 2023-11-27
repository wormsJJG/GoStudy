package main

import "fmt"

func status(name string) {

	balls := map[string]int{"성기훈": 20, "오일남": 0}

	ball, exists := balls[name]
	fmt.Println(ball, exists)

	if !exists {

		fmt.Println(name, "님은 게임 참여자가 아닙니다.")
	} else if ball < 1 {

		fmt.Println(name, "님은", ball, "개로 탈락")
	} else {

		fmt.Println(name, "님은 게임에서 승리하셨습니다!")
	}
}

func main() {

	status("오일남")
	status("강철")
	status("성기훈")
	// var balls map[string]int
	// fmt.Println(balls)
	// balls := make(map[string]int)
	// fmt.Printf("%#v\n", balls)
	// balls["성기훈"] = 20
	// fmt.Println(balls["성기훈"])
	// fmt.Println(balls["오일남"])
}

// package main

// import "fmt"

// func main() {

// 	// var games map[int]string
// 	// games := make(map[int]string)
// 	games := map[int]string{
// 		456: "성기훈",
// 		218: "박해수",
// 		067: "강새벽",
// 		001: "오일남",
// 		199: "오일남",
// 		101: "아이오아이",
// 	}

// 	// games[456] = "성기훈"
// 	// games[218] = "박해수"
// 	// games[067] = "강새벽"
// 	// games[001] = "오일남"
// 	// games[199] = "알리"

// 	for _, value := range games {

// 		fmt.Println(value)
// 	}

// 	games[101] = "장덕수"
// 	delete(games, 199)

// 	for k, v := range games {

// 		fmt.Println(k, v)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/headfirstgo/datafile"
// )

// func main() {

// 	lines, err := datafile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	counts := make(map[string]int)
// 	for _, line := range lines {
// 		counts[line]++
// 	}
// 	for name, count := range counts {
// 		fmt.Printf("Votes for %s: %d\n", name, count)
// 	}
// }
