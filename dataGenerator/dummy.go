package dataGenerator

import (
	"fmt"
	"github.com/manveru/faker"
	"strings"
)

func Dummy() {
	fake, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	title := strings.ToUpper(fake.NamePrefix())
	name := strings.ToUpper(fake.Name())

	fmt.Println("NAME:", title, name)
	fmt.Println("EMAIL:", fake.Email())
	fmt.Println("ADDRESS:", fake.StreetAddress(), fake.City(), fake.PostCode(), fake.State(), fake.Country())
	fmt.Println("PHONE NUMBER:", fake.PhoneNumber())
	fmt.Println("MOBILE:", fake.CellPhoneNumber())
	fmt.Println("-----------------------", "\n")
}