package mymain

import (
	"fmt"

	"github.com/dongri/phonenumber"
)

func mymain() {
	// Get country with mobile and land line numbers
// Let's try to get country for Latvian land line number
includeLandLine := true
country := phonenumber.GetISO3166ByNumber("081367092298", includeLandLine)
fmt.Println(country.CountryName)
// Output: Latvia
}