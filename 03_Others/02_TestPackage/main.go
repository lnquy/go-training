package main

import (
	"fmt"
	"gitlab.com/lnquy/go-training/03_Others/02_TestPackage/pkg1"
)

func main() {
	fmt.Printf("The exported var: %d", mypkg.ExportedVar)
}
