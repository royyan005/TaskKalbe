package mymain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



var truephone string ="6281367092298"

func TestPhoneIsValid(t *testing.T) {
	var num1 = PhoneIsValid("081367092298") // diawali 0
	assert.Equal(t, num1, truephone, "phone not equal")

	var num2 = PhoneIsValid("6281367092298") // normal
	assert.Equal(t, num2, truephone, "phone not equal")

	var num3 = PhoneIsValid("+6281367092298") // diawali +
	assert.Equal(t, num3, truephone, "phone not equal")

	var num4 = PhoneIsValid("08136709229800000") // lebih dari 16 digit
	assert.Equal(t, num4, "ERROR!", "phone not equal")
}

