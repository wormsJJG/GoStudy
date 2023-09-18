package main

import (
	"fmt"
)

func main() {

	price := 100
	fmt.Println("Price is", "price", "dollars.")

	taxPate := 0.08
	tax := float64(price) * taxPate
	fmt.Println("Tax is", tax, "dollars.")

	total := price + int(tax)
	fmt.Println("Totla cost is", total, "dollars.")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars acailable")
	fmt.Println("Within budget?", availableFunds > int(total))
}
