package gopackexample

import "fmt"

func PackageVersion() {
	fmt.Println("Package version - 1.1.0")
}

func PackageFunctionality(a, b int) int {
	return a + b
}
