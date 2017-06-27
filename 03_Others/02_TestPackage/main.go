package main

import (
	"gitlab.com/lnquy/go-training/03_Others/02_TestPackage/pkg1"
	"fmt"
)

func main() {
	fmt.Printf("The exported var: %d", mypkg.ExportedVar)
}
