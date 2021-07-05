package gopackexample

import "fmt"

func PackageVersion() {
	fmt.Println("Package version - 2.1.0")
}

func PackageFunctionality(a, b int) int {
	return a * b
}

func PackageFunctionality2(a, b, c int) int {
	return a + b*c
}
