package gopackexample

import (
	"fmt"
	"math"
)

func PackageVersion() {
	fmt.Println("Package version - 3.0.0")
}

func PackageFunctionality(a, b float64) float64 {
	return math.Pow(a, b)
}

func PackageFunctionality2(a, b, c float64) float64 {
	return a * math.Pow(b, c)
}
