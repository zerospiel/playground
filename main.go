package main

import (
	"cmp"
	"fmt"
	"iter"
	"os"
	"runtime/trace"
	"strconv"
)

func main() {
}

func FooIter[E any](s []E) iter.Seq2[int, E] {
	defer func() {
		println("immediately after the evaluation of the for loop")
	}()
	return func(yield func(int, E) bool) {
		defer func() {
			if rc := recover(); rc != nil {
				fmt.Println("panic:", rc)
			}
		}()
		for i := 0; i < len(s); i++ {
			defer func() {
				println("iter #" + strconv.Itoa(i))
			}()
			if !yield(i, s[i]) {
				return
			}
		}
	}
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
