package main

import (
	"fmt"
	"github.com/TatianaThorpe/learnGoProject/dataGenerator"
)

func main() {

	var number int
	fmt.Print("Please enter the number of records you require: ")
	fmt.Scan(&number)
	i := 0

	for i < number {
		i++
		println("\n---------- RECORD", i, "----------")
		dataGenerator.Dummy()
	}

	fmt.Println("\nRECORDS GENERATED:", number)
}
