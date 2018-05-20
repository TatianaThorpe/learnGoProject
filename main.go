package main

import (
	"fmt"
	"github.com/manveru/faker"
	"strings"
	"os"
)

var (
	min = 1
	max = 10000
)

func main() {
	number := userinput()
	i := 0
	countryMatch := 0

	for i < number {
		i++
		println("\n---------- RECORD", i, "----------")
		fake, err := faker.New("en")
		if err != nil {
			panic(err)
		}

		title := strings.ToUpper(fake.NamePrefix())
		name := strings.ToUpper(fake.Name())
		country := fake.Country()

		fmt.Println("NAME:", title, name)
		fmt.Println("EMAIL:", fake.Email())
		fmt.Println("ADDRESS:", fake.StreetAddress(), fake.City(), fake.PostCode(), fake.State(), country)
		fmt.Println("PHONE NUMBER:", fake.PhoneNumber())
		fmt.Println("MOBILE:", fake.CellPhoneNumber())

		if country == "United Kingdom" || country == "United States of America" {
			countryMatch++
		}

	}

	fmt.Println("\nRECORDS GENERATED:", number)
	fmt.Println("Number of UK or USA records found:", countryMatch)
}

func userinput() int {
	var number int

	fmt.Printf("Please enter a number between %d and %d: ", min, max)
	fmt.Scan(&number)

	if number < min || number > max {
		fmt.Printf("Please enter a number between %d and %d\n", min, max)
		os.Exit(1)
	}
	return number
}
