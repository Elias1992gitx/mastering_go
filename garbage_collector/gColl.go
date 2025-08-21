//The purpose of the printStats() function is to avoid writing the same Go code all of the time.

//The operation of the Go garbage collector is based on the tricolor algorithm

package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Print("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {

	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)

		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	
	printStats(mem)


	for i := 0; i < 10; i++ {
		s := make ([]byte, 100000000)

		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep(5 * time.Second)
	}

	printStats(mem)

}



/*More about the operation of the Go Garbage
Collector - page 53*/
