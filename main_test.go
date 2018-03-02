package main

import "testing"

const size = 100000

func swapAppend(s *[]int) {
	l := len((*s)) - 1
	(*s) = append([]int{(*s)[l]}, (*s)[:l]...)
}

func swapAppendPreallocate(s *[]int) {
	l := len((*s))
	temp := make([]int, 1, l)
	l = l - 1
	temp[0] = (*s)[l]
	(*s) = append(temp, (*s)[:l]...)
}

func swapCopy(s []int) {
	n := s[len(s)-1]
	copy(s[1:], s[0:len(s)-1])
	s[0] = n
}

func BenchmarkCopy(b *testing.B) {
	a := make([]int, size, size)
	for i := 0; i < b.N; i++ {
		swapCopy(a)
	}
}

func BenchmarkAppend(b *testing.B) {
	a := make([]int, size, size)
	for i := 0; i < b.N; i++ {
		swapAppend(&a)
	}
}

func BenchmarkAppendPreallocate(b *testing.B) {
	a := make([]int, size, size)
	for i := 0; i < b.N; i++ {
		swapAppendPreallocate(&a)
	}
}
