package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	fmt.Printf("runtime.Version(): %v\n", runtime.Version())
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
