package main

import "fmt"

const UsdToEur = 0.93
const UsdToRub = 88.2

func main() {
	eurToRub := UsdToRub / UsdToEur
	fmt.Printf("Hello! %v", eurToRub)
}
