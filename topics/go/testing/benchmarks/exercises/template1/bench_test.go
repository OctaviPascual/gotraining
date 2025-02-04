// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -run none -bench . -benchtime 3s -benchmem

// Write three benchmark tests for converting an integer into a string. First using the
// fmt.Sprintf function, then the strconv.FormatInt function and then strconv.Itoa.
// Identify which function performs the best.
package main

import (
	"fmt"
	"strconv"
	"testing"
)

const number = 12341234

var gs string

// BenchmarkSprintf provides performance numbers for the fmt.Sprintf function.
func BenchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%d", number)
	}

	gs = s
}

// BenchmarkFormat provides performance numbers for the strconv.FormatInt function.
func BenchmarkFormat(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = strconv.FormatInt(number, 10)
	}

	gs = s
}

// BenchmarkItoa provides performance numbers for the strconv.Itoa function.
func BenchmarkItoa(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = strconv.Itoa(number)
	}

	gs = s
}
