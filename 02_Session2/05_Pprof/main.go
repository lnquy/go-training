package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Before poorFunc")
	poorFunc(5)
	fmt.Println("After poorFunc")

	for i := 0; i < 100000000; i++ {

	}
}

func poorFunc(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Goroutine was panicked (%s)! But we can recover from it :))\n", err)
		} else {
			fmt.Println("Goroutine is ok!")
		}
	}()

	if 5/i == 0 {
		fmt.Println("Impossible")
	} else {
		fmt.Println("Possible")
	}
}

// go tool pprof 05_Pprof.exe aaa.prof
