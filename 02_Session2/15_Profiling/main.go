package main

import (
	"fmt"
	"time"
	//_ "net/http/pprof"
	//"net/http"
	//"log"
)

func checkPrime1(n int) (isPrime bool) {
	defer execTime("checkPrime1", time.Now())

	for i := n-1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func execTime(name string, t time.Time) {
	fmt.Printf("%v executed in: %v\n", name, time.Now().Sub(t))
}

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:15000", nil))
	//}()

	nu := 982451653
	fmt.Printf("%v is prime: %v\n", nu, checkPrime1(nu))
	fmt.Scanln()
}
