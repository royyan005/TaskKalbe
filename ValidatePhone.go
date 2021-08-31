package mymain

import (
	// "fmt"
	"regexp"
	"strings"
)

// func main() {
// 	fmt.Println(PhoneIsValid("081367092298"))
// 	fmt.Println(PhoneIsValid("6281367092298"))
// 	fmt.Println(PhoneIsValid("+6281367092298"))
// 	fmt.Println(PhoneIsValid("08136709229800000"))
// 	fmt.Println(PhoneIsValid("081367asdaaw"))
// 	fmt.Println(PhoneIsValid("08136"))
// }

var Email string
var PhoneNumber string

func PhoneIsValid(PhoneNumber string) string {
	// parseall := regexp.MustCompile(`^(\+62|62|0)8[1-9][0-9]{6,9}$`)
	parseplus := regexp.MustCompile(`^(\+62)8[1-9][0-9]{6,9}$`)
	parse62 := regexp.MustCompile(`^(62)8[1-9][0-9]{6,9}$`)
	parse08 := regexp.MustCompile(`^(0)8[1-9][0-9]{6,9}$`)

	if parseplus.MatchString(PhoneNumber) {
		return strings.Trim(PhoneNumber, "+")
	} else if parse62.MatchString(PhoneNumber) {
		return PhoneNumber
	} else if parse08.MatchString(PhoneNumber) {
		return "62" + strings.Trim(PhoneNumber, "0")
	} else if strings.Count(PhoneNumber, "") > 16 {
		return "ERROR!"
	} else if strings.Count(PhoneNumber, "") < 11 {
		return "ERROR!"
	}
	return "ERROR!"
}
