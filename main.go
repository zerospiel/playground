package main

import (
	"constraints"
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/trace"
	"sort"
	"strconv"
	"strings"

	generatedv1 "github.com/zerospiel/playground/gen/go/pg/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//go:generate go run google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 --version
func noop() {}

type foobar struct {
	generatedv1.UnimplementedStringsServiceServer
}

func (*foobar) ToUpper(ctx context.Context, req *generatedv1.ToUpperRequest) (*generatedv1.ToUpperResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(metadata.FromIncomingContext(ctx))
	fmt.Println(metadata.FromOutgoingContext(ctx))
	return &generatedv1.ToUpperResponse{
		S: strings.ToUpper(req.S),
	}, nil
}

type sss struct{}

func (*sss) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bb, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(bb))
	w.WriteHeader(200)
	return
}

func map2slices[T comparable, V any](m map[T]V) ([]T, []V) {
	ks, vs := make([]T, 0, len(m)), make([]V, 0, len(m))
	for k, v := range m {
		ks, vs = append(ks, k), append(vs, v)
	}
	return ks, vs
}

type nestedT struct {
	intptr *int64
	sptr   *(struct {
		name     string
		lastname string
	})
}

type RPValue interface {
	constraints.Integer | constraints.Float
}

type ResolverParams[K comparable, V RPValue] map[K]V

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

type node struct {
	next *node
	data any
}

func hasLoop(list *node) bool {
	if list == nil || list.next == nil {
		return false
	}

	slow, fast := list, list
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}

	return false
}

func push(head **node, data any) {
	newN := new(node)
	newN.data = data
	newN.next = *head
	(*head) = newN
}

func reverse(head **node) {
	prev, cur := new(node), *head
	for cur != nil {
		next := cur.next

		cur.next = prev
		prev = cur
		cur = next
	}
	*head = prev
}

func fib(n int) int {
	var f = make([]int, n)
	f[0], f[1] = 1, 1
	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n-1]
}

func printlist(list *node) {
	for list != nil {
		if list.next == nil {
			fmt.Printf("%v", list.data)
		} else {
			fmt.Printf("%v -> ", list.data)
		}
		list = list.next
	}
	fmt.Println()
}

func binsearch(xs []int, target int) int {
	l, r := 0, len(xs)-1
	for l <= r {
		mid := l + (r-l)/2
		if xs[mid] == target {
			return mid
		}
		if xs[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func binsearch2(xs []int, target int) int {
	l, r := 0, len(xs)-1
	for r-l > 1 {
		mid := l + (r-l)/2
		if xs[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if xs[l] == target {
		return l
	}
	if xs[r] == target {
		return r
	}
	return -1
}

func LIS(nums []int) []int {
	lis := make([][]int, len(nums))
	lis[0] = append(lis[0], nums[0])

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && len(lis[i]) < len(lis[j])+1 {
				lis[i] = lis[j]
			}
		}
		lis[i] = append(lis[i], nums[i])
	}

	maximum := lis[0]
	for i := 0; i < len(lis); i++ {
		if len(lis[i]) > len(maximum) {
			maximum = lis[i]
		}
	}

	return maximum
}

func isPowOfTwo(target int) bool {
	ones := 0
	for target > 0 {
		if target&1 == 1 {
			ones++
		}
		target >>= 1
	}
	return ones == 1
}

func median(nums []int) float64 {
	n := len(nums)
	if n&1 != 0 {
		return float64(nums[n/2])
	}
	return float64((nums[n/2-1] + nums[n/2])) / 2.
}

func medianArrs(a, b []int) float64 {
	if len(a) == 0 {
		return median(b)
	}
	if len(b) == 0 {
		return median(a)
	}
	if len(a) == 0 && len(b) == 0 {
		return 0
	}

	small, larger := a, b // aliasing
	if len(a) > len(b) {
		small, larger = b, a
	}

	n, m := len(small), len(larger)
	left, right, mergedMid := 0, n, (n+m+1)/2

	for left <= right {
		mid := left + (right-left)/2
		leftSmallSize, leftLargeSize := mid, mergedMid-mid

		var (
			leftSmall, leftLarge   = math.MinInt, math.MinInt
			rightSmall, rightLarge = math.MaxInt, math.MaxInt
		)

		if leftSmallSize > 0 {
			leftSmall = small[leftSmallSize-1]
		}
		if leftSmallSize < n {
			rightSmall = small[leftSmallSize]
		}

		if leftLargeSize > 0 {
			leftLarge = larger[leftLargeSize-1]
		}
		if leftLargeSize < m {
			rightLarge = larger[leftLargeSize]
		}

		if leftSmall <= rightLarge && leftLarge <= rightSmall {
			if (m+n)&1 != 0 {
				return math.Max(float64(leftSmall), float64(leftLarge))
			}

			return (math.Min(float64(rightSmall), float64(rightLarge)) + math.Max(float64(leftSmall), float64(leftLarge))) / 2.
		} else if leftSmall > leftLarge {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return 0
}

func medianTwoSorted(a, b []int) float64 {
	i, j, mergedLen := 0, 0, len(a)+len(b)
	pm, m := 0., 0.
	for k := 0; k <= mergedLen/2; k++ {
		if mergedLen&1 == 0 {
			pm = m
		} else {
			pm = 0.
		}

		if i != len(a) && j != len(b) {
			if a[i] < b[j] {
				m = float64(a[i])
				i++
			} else {
				m = float64(b[j])
				j++
			}
		} else {
			if i < len(a) {
				m = float64(a[i])
				i++
			} else {
				m = float64(b[j])
				j++
			}
		}
	}
	if mergedLen&1 == 0 {
		return (m + pm) / 2
	}
	return m
}

func supressString(s string) string {
	if !(s[0] >= 'A' && s[0] <= 'Z') {
		return ""
	}

	count := 1
	var result strings.Builder
	last := s[0]
	for i := 1; i < len(s); i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') {
			return ""
		}
		if s[i-1] == s[i] {
			count++
		} else {
			if count == 1 {
				result.WriteByte(last)
			} else {
				result.WriteByte(last)
				result.WriteString(strconv.Itoa(count))
			}
			last = s[i]
			count = 1
		}
	}

	if count == 1 {
		result.WriteByte(s[len(s)-1])
	} else {
		result.WriteByte(s[len(s)-1])
		result.WriteString(strconv.Itoa(count))
	}
	return result.String()
}

func intersectWithRepeats(a, b []int) []int {
	m := make(map[int]int)
	for _, v := range b {
		m[v]++
	}
	res := []int{}
	for _, v := range a {
		if c, ok := m[v]; ok && c > 0 {
			res = append(res, v)
		}
	}
	return res
}

func getSegments(nums []int) string {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	result := ""
	groupStart := nums[0]
	groupEnd := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			groupEnd = nums[i]
		} else {
			if groupEnd == groupStart {
				result += strconv.Itoa(groupStart) + ","
			} else {
				result += fmt.Sprintf("%d-%d,", groupStart, groupEnd)
			}
			groupEnd = nums[i]
			groupStart = nums[i]
		}
	}
	if groupEnd == groupStart {
		result += strconv.Itoa(groupStart) + ","
	} else {
		result += fmt.Sprintf("%d-%d,", groupStart, groupEnd)
	}
	return result[:len(result)-1]
}

func maxOnesSubs(arr []int) int {
	pre, post := make([]int, len(arr)), make([]int, len(arr))
	pre[0], post[len(arr)-1] = 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == 1 && arr[i-1] == 1 {
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = 1
		}
	}
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] == 1 && arr[i+1] == 1 {
			post[i] = post[i+1] + 1
		} else {
			post[i] = 1
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := 0
	length := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == 1 && arr[i-1] == 1 {
			length++
		} else {
			length = 1
		}
		result = max(result, length)
	}
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == 1 && arr[i+1] == 1 {
			result = max(result, pre[i-1]+post[i+1])
		}
	}
	return result
}

func main() {
	println(43%6, 42%6, 5%6)
	return

	fmt.Println(medianArrs([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5}))
	fmt.Println(medianArrs([]int{2, 3, 5, 8}, []int{10, 12, 14, 16, 18, 20}))
	fmt.Println(medianArrs([]int{-5, 3, 6, 12, 15}, []int{-12, -10, -6, -3, 4, 10}))
	fmt.Println(medianTwoSorted([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5}))
	fmt.Println(medianTwoSorted([]int{2, 3, 5, 8}, []int{10, 12, 14, 16, 18, 20}))
	fmt.Println(medianTwoSorted([]int{-5, 3, 6, 12, 15}, []int{-12, -10, -6, -3, 4, 10}))
	return

	println(fib(34))

	fmt.Println(LIS([]int{3, 2, 6, 4, 5, 1}))

	return

	head := new(node)
	push(&head, 20)
	push(&head, 4)
	push(&head, 15)
	push(&head, 10)
	// printlist(head)
	// reverse(&head)
	printlist(head)
	head.next.next.next = head
	println(hasLoop(head))
	println(binsearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	println(binsearch2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))

	return
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
