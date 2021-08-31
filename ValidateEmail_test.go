package mymain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



var (
	truemail bool = true
)

func TestIsEmailValid(t *testing.T) {
	testemail1 := isEmailValid("test44@gmail.com")
	assert.Equal(t, testemail1, true, "Wrong Email Format!")

	testemail2 := isEmailValid("tesT44@gmail.com")
	assert.Equal(t, testemail2, true, "Wrong Email Format!")

	testemail3 := isEmailValid("tesT44@gmail")
	assert.Equal(t, testemail3, false, "Wrong Email Format!")

	testemail4 := isEmailValid("@gmail.com")
	assert.Equal(t, testemail4, false, "Wrong Email Format!")
}

