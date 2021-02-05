package main

import (
  "fmt"
  "math"  // imported math function to demonstrate exponents
)

func main() {
  fmt.Println("Addition:", 1 + 3)
  fmt.Println("Subtraction:", 1 - 3)
  fmt.Println("Multiplication:", 2 * 3)
  fmt.Println("Division:", 20 / 4)  // Division without a float
  fmt.Println("Division:", 20.0 / 3)  // Division with a float.  adding the decimal
  fmt.Println("Exponents:", math.Pow(20.0, 3))

}

