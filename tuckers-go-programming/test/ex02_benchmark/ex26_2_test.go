package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci1(t *testing.T) {
	assert := assert2.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci(-1) = 0")
	assert.Equal(0, fibonacci1(0), "fibonacci(0) = 0")
	assert.Equal(1, fibonacci1(1), "fibonacci(1) = 1")
	assert.Equal(2, fibonacci1(3), "fibonacci(2) = 2")
	assert.Equal(233, fibonacci1(13), "fibonacci(13) = 233")
}
func TestFibonacci2(t *testing.T) {
	assert := assert2.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci(-1) = 0")
	assert.Equal(0, fibonacci2(0), "fibonacci(0) = 0")
	assert.Equal(1, fibonacci2(1), "fibonacci(1) = 1")
	assert.Equal(2, fibonacci2(3), "fibonacci(3) = 2")
	assert.Equal(233, fibonacci2(13), "fibonacci(13) = 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
