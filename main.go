package main

import (
	"fmt"
	"github.com/TatianaThorpe/learnGoProject/dataGenerator"
	//"go/types"
)

func main() {

	var number int
	fmt.Print("Please enter the number of records you require: ")
	fmt.Scanf("%d", &number)
	i := 0

	for i < number {
		i++
		println("\n---------- RECORD", i,"----------")
		dataGenerator.Dummy()
	}

	fmt.Println("\nRECORDS GENERATED:", number)
}
