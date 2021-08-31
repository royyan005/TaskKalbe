package mymain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



var truephone string = PhoneIsValid("081367092298")

func TestPhoneIsValid(t *testing.T) {
	var num = PhoneIsValid("081367092298")
	assert.Equal(t, num, truephone, "phone equal")
}

