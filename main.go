package main

import (
	"cmp"
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	m := map[string]float64{"1": 1., "2": 2., "0.2": .2}
	fmt.Printf("M2S[[]float64](m): %v\n", M2S[[]float64](m))
}

func M2S[S ~[]V, M ~map[K]V, K comparable, V cmp.Ordered](m M) S {
	s := make(S, 0)
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func getTrace() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	trace.Stop()
}
