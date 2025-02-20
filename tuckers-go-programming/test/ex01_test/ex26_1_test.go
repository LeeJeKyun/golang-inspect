package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare3(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(81, square(9), "9*9 == 81")
}

func TestSquare4(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(9, square(3), "3*3 == 9")
}
