package gopackexample

import (
	"fmt"
	"math"
)

func PackageVersion() {
	fmt.Println("Package version - 3.1.0")
}

func PackageFunctionality(a, b float64) float64 {
	return math.Pow(a, b)
}

func PackageFunctionality2(a, b, c float64) float64 {
	return a * math.Pow(b, c)
}

func PackageFunctionality3(a, b, c, x float64) float64 {
	return a*math.Pow(x, 2) + b*x + c
}
