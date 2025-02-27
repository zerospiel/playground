package main

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"fmt"
	"hash/maphash"
	"iter"
	"maps"
	"os"
	"runtime"
	"runtime/trace"
	"slices"
	"strconv"
	"unique"
	"weak"
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

type TreeIntStr[K int, V comparable] = Tree[K, V] // made-up dummy example, GOEXPERIMENT=aliastypeparams

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

func sortfn() {
	const cnt = 20
	m := make(map[int]struct{}, cnt)
	for i := range cnt {
		m[i] = struct{}{}
	}

	keys := maps.Keys(m)
	sortedKeys := slices.Sorted(keys)
	unsortedKeys := slices.Collect(keys)
	fmt.Printf("sorted: %v; unsorted: %v\n", sortedKeys, unsortedKeys)
}

func mhash() {
	s := maphash.MakeSeed()
	v1, v2 := unique.Make(100), unique.Make(101)
	h1 := maphash.Comparable(s, v1)
	h2 := maphash.Comparable(s, v2)
	fmt.Printf("h1: %v\n", h1)
	fmt.Printf("h2: %v\n", h2)
	fmt.Printf("(h1 == h2): %v\n", (h1 == h2))
	v3, v4 := 100, 100
	h1, h2 = maphash.Comparable(s, v3), maphash.Comparable(s, v4)
	fmt.Printf("h1: %v\n", h1)
	fmt.Printf("h2: %v\n", h2)
	fmt.Printf("(h1 == h2): %v\n", (h1 == h2))
}

func wp() {
	s := new(string)
	println("original:", s)

	wp := weak.Make(s)
	runtime.GC() // wp underlaying ptr still needed bc of the next println usage

	ptr := wp.Value()
	println("value before gc:", ptr, s)
	runtime.GC()

	ptr = wp.Value()
	println("value after gc:", ptr)
}

type TZero struct {
	A myInt `json:"a,omitzero"`
}

type myInt int

func (t myInt) IsZero() bool {
	return t < 1
}

func ejson() {
	t := &TZero{A: -1}
	bb, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("string(bb): %v\n", string(bb))
}

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
