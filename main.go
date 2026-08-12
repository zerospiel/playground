package main

import (
	// _ "simd/archsimd"
	"cmp"
	"context"
	_ "embed"
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
	"hash/maphash"
	"iter"
	"maps"
	"net/netip"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	// "simd" // GOEXPERIMENT=simd
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unique"
	"uuid"
	"weak"
)

func main() {
}

// 1.27

// generic methods

type Box[T any] struct{ v T }

func (b Box[T]) Map[U any](f func(T) U) Box[U] {
	return Box[U]{v: f(b.v)}
}

func genericMethod() {
	b := Box[int]{v: 21}
	doubled := b.Map(func(n int) int { return n * 2 })
	label := doubled.Map(func(n int) string {
		return fmt.Sprintf("value=%d", n)
	})
	fmt.Println(label.v)
}

// struct literal field selectors

func structLiteralFieldSelectors() {
	type (
		Base struct {
			ID int
		}

		User struct {
			Name string
			Base
		}
	)

	u := User{ID: 7, Name: "Mittens"}
	fmt.Println(u.ID, u.Name)
}

// generalized function type inference

func first[T any](s []T) T { return s[0] }
func last[T any](s []T) T  { return s[len(s)-1] }
func fnTypeInference() {
	// The slice's element type drives inference: T=int for each entry.
	// Before Go 1.27 this failed with "cannot use generic function
	// without instantiation"; you had to write first[int], last[int].
	ops := []func([]int) int{first, last}
	for _, op := range ops {
		fmt.Println(op([]int{10, 20, 30}))
	}
}

// labels in tracebacks

func labeledTracebacks() {
	ctx := context.Background()
	pprof.Do(ctx, pprof.Labels("request", "42"), func(ctx context.Context) {
		buf := make([]byte, 1<<12)
		n := runtime.Stack(buf, false)
		fmt.Printf("%s", buf[:n])
	})
}

// goroutineleak profile graduated

func leakProfileGraduated() {
	leak := func() {
		ch := make(chan int)
		ch <- 1
	}

	go leak()

	runtime.Gosched()

	// The GC-backed scan finds goroutines that can never make progress.
	pprof.Lookup("goroutineleak").WriteTo(os.Stdout, 1)
}

// uuid pkg

func uuidPkg() {
	a := uuid.MustParse("f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	fmt.Println("parsed:", a)
	fmt.Println("nil:   ", uuid.Nil())
	fmt.Println("max:   ", uuid.Max())
	fmt.Println(uuid.NewV4()) // random
	fmt.Println(uuid.NewV7()) // time-ordered
}

// non-fixed (size-agnostic) SIMD

// func sizeAgnosticSIMD() {
// 	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
// 	b := []float32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160}

// 	va := simd.LoadFloat32s(a) // reads exactly va.Len() lanes from a
// 	vb := simd.LoadFloat32s(b)

// 	sum := va.Add(vb) // element-wise add, many lanes in one instruction

// 	out := make([]float32, sum.Len())
// 	sum.Store(out)

// 	fmt.Println("simd:", out)
// }

// cutLast

func cutLast() {
	before, after, found := strings.CutLast("a/b/c", "/")
	fmt.Printf("%q %q %v\n", before, after, found)

	before, after, found = strings.CutLast("nosep", "/")
	fmt.Printf("%q %q %v\n", before, after, found)
}

// generic hash

type ciHasher struct{}

// Equal ignores case; Hash mixes in the lower-cased form, so values
// that are Equal always hash the same.
func (ciHasher) Hash(h *maphash.Hash, s string) { h.WriteString(strings.ToLower(s)) }
func (ciHasher) Equal(x, y string) bool         { return strings.EqualFold(x, y) }

func genericHash() {
	var h maphash.Hasher[string] = ciHasher{} // plug in the custom strategy

	fmt.Println("generic hash pure equal:", h.Equal("Go", "GO"), h.Equal("Go", "Rust"))

	// Equal values must hash the same, so feed each into a Hash sharing one seed:
	seed := maphash.MakeSeed()
	var a, b maphash.Hash
	a.SetSeed(seed)
	b.SetSeed(seed)
	h.Hash(&a, "Go")
	h.Hash(&b, "GO")
	fmt.Println("generic hash sums equality:", a.Sum64() == b.Sum64())
}

// 1.26

func newexpr() {
	pi, ps, pf := new(42), new([]string{"42"}), new(func() { _ = 42 })

	fmt.Printf("newexpr type: %T; %T; %T\n", pi, ps, pf)
}

func cidrComp() {
	cidr1 := netip.MustParsePrefix("10.0.0.0/16")
	cidr2 := netip.MustParsePrefix("10.0.0.0/15")
	fmt.Printf("cidr1: %v\n", cidr1)
	fmt.Printf("cidr2: %v\n", cidr2)
	fmt.Printf("cidr1.Compare(cidr2): %v\n", cidr1.Compare(cidr2))
}

func signalCause() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	p, _ := os.FindProcess(os.Getpid())
	_ = p.Signal(syscall.SIGINT)

	<-ctx.Done()
	fmt.Println("err =", ctx.Err())
	fmt.Println("cause =", context.Cause(ctx))
}

func recTypes() {
	type Ordered[T Ordered[T]] interface {
		Less(T) bool
	}
	type Tree[T Ordered[T]] struct {
		nodes []T
	}

	t := Tree[netip.Addr]{}
	_ = t
}

func leakProfile() {
	gather := func(funcs ...func() int) <-chan int {
		out := make(chan int)
		for _, f := range funcs {
			go func() {
				out <- f()
			}()
		}
		return out
	}

	printLeaks := func(f func()) {
		// not required since 1.27
		// if !strings.Contains(os.Getenv("GOEXPERIMENT"), "goroutineleakprofile") {
		// 	panic("set GOEXPERIMENT=goroutineleakprofile")
		// }

		prof := pprof.Lookup("goroutineleak")

		defer func() {
			time.Sleep(50 * time.Millisecond)
			var content strings.Builder
			prof.WriteTo(&content, 2)
			// Print only the leaked goroutines.
			for goro := range strings.SplitSeq(content.String(), "\n\n") {
				if strings.Contains(goro, "(leaked)") {
					fmt.Println(goro + "\n")
				}
			}
		}()

		f()
	}
	printLeaks(func() {
		gather(
			func() int { return 11 },
			func() int { return 22 },
			func() int { return 33 },
		)
	})
}

// 1.25

func callJsonv2() {
	// not required since 1.27
	// if !strings.Contains(os.Getenv("GOEXPERIMENT"), "jsonv2") {
	// 	panic("set GOEXPERIMENT=jsonv2")
	// }

	boolMarshaler := jsonv2.MarshalFunc(
		func(val bool) ([]byte, error) {
			if val {
				return []byte(`"✅"`), nil
			}
			return []byte(`"🔪"`), nil
		},
	)

	val := false
	bb, err := jsonv2.Marshal(val, jsonv2.WithMarshalers(boolMarshaler))
	if err != nil {
		panic(err)
	}

	fmt.Printf("string(bb): %v\n", string(bb))
}

// 1.24

func aliasTypeParams() {
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
}

func maphashCompare() {
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

type TZero struct {
	A myInt `json:"a,omitzero"`
}

type myInt int

func (t myInt) IsZero() bool {
	return t < 1
}

func jsonOmitZero() {
	t := &TZero{A: -1}
	bb, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("string(bb): %v\n", string(bb))
}

func weakpointer() {
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

// 1.23 and older (1.22-)

func seqsSort() {
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

// 1.18

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
