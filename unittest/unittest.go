package main

import "gitlab.com/lnquy/go-training/unittest/foo3"

func main() {
	foo3Reader := foo3.Reader{}
	foo3Reader.Fetch("") // Run this main() to check what GetData() will be called
}
