package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	go gocount()
	Process.Init()
	for {
		Process.Mem.Update()
		Process.Cpu.Update()
		fmt.Println(Process.Cpu, Process.Mem)
		time.Sleep(1 * time.Second)
	}
}

var xxx int64

func gocount() {
	for {
		xxx++
		time.Sleep(1e4)
	}
}
