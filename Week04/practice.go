package main

import (
	"fmt"
)

func main() {

	price := 100 //int
	fmt.Println("Price is", "price", "dollars.")

	taxPate := 0.08 //float64
	tax := float64(price) * taxPate //float64
	fmt.Println("Tax is", tax, "dollars.")

	total := price + int(tax)
	fmt.Println("Total cost is", total, "dollars.")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars acailable")
	fmt.Println("Within budget?", availableFunds > int(total))
}
