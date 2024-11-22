package main

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"os"
	"runtime/trace"
	"slices"
	"strconv"
)

// type foo struct {
// 	_ structs.HostLayout
// }

type Tree[K cmp.Ordered, V any] struct {
	left, right *Tree[K, V]
	key         K
	value       V
}

type TreeInt = Tree[int, any]

// type TreeIntStr[K int, V string] = Tree[K, V] // made-up dummy example, GOEXPERIMENT=aliastypeparams

func panicIter() {
	defer func() {
		if p := recover(); p != nil {
			println("main panic:", p)
			panic(p)
		}
	}()

	next, _ := iter.Pull(func(yield func(V any) bool) {
		yield("hello")
		panic("world")
	})

	for {
		fmt.Println(next())
	}
}

func main() {
	const cnt = 20
	m := make(map[int]struct{}, cnt)
	for i := range cnt {
		m[i] = struct{}{}
	}

	keys := maps.Keys(m)
	sortedKeys := slices.Sorted(keys)
	unsortedKeys := slices.Collect(keys)
	fmt.Printf("sorted: %v; unsorted: %v\n", sortedKeys, unsortedKeys)

	// panicIter()
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
