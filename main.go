package main

import (
	"fmt"
	"github.com/manveru/faker"
	"strings"
)

func main() {
	fake, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	var number int
	fmt.Print("Please enter the number of records you require: ")
	fmt.Scan(&number)
	i := 0

	for i < number {
		i++
		title := strings.ToUpper(fake.NamePrefix())
		name := strings.ToUpper(fake.Name())

		println("START OF RECORD", i)
		println("NAME:", title, name)
		println("EMAIL:", fake.Email())
		println("ADDRESS:", fake.StreetAddress(), fake.City(), fake.PostCode(), fake.State(), fake.Country())
		println("PHONE NUMBER:", fake.PhoneNumber())
		println("MOBILE:", fake.CellPhoneNumber(), "\n")
	}

	fmt.Println("RECORDS GENERATED:", number)
}
