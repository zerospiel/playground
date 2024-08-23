package main

import (
	"iter"
	"maps"
	"slices"
	"strconv"
	"testing"
)

func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func LongStrings(m map[int]string, n int) []string {
	isLong := func(s string) bool {
		return len(s) >= n
	}
	return slices.Collect(Filter(isLong, maps.Values(m)))
}

func LongStringsSimple(m map[int]string, n int) []string {
	s := make([]string, 0)
	for _, v := range m {
		if len(v) >= n {
			s = append(s, v)
		}
	}

	return s
}

func getMap(n int) map[int]string {
	m := make(map[int]string, n)
	for i := range n {
		m[i] = strconv.Itoa(i)
	}
	return m
}

const (
	Nm int = 1e6
	Nc     = 5
)

func BenchmarkFilter(b *testing.B) {
	m := getMap(Nm)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = LongStrings(m, Nc)
	}
}

func BenchmarkSimple(b *testing.B) {
	m := getMap(Nm)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = LongStringsSimple(m, Nc)
	}
}
