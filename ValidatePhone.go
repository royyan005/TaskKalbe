package mymain

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(PhoneIsValid("081367092298"))
}

	var Email string
	var PhoneNumber  string


func PhoneIsValid(PhoneNumber string) string {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(PhoneNumber) {
		return PhoneNumber
	}
	return PhoneNumber
}
