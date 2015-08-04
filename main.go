package main

import (
	pp "./v1"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pp.Run()
}
