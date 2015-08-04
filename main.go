package main

import (
	pp "go.iondynamics.net/passPad/v1"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pp.Run()
}
