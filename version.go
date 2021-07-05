package gopackexample

import "fmt"

func PackageVersion() {
	fmt.Println("Package version - 2.0.0")
}

func PackageFunctionality(a, b int) int {
	return a * b
}
