package main

import (
	"fmt"
	"runtime"
)

func main() {
	b := runtime.NumCPU()
	fmt.Println(b)
}