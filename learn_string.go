package main

import "fmt"

func main() {
  fmt.Println("Simple String")
  fmt.Println(`
  This is a multi-line. \n
  String
  `)
  fmt.Println(`a string with back tick instead of single quotes`)
  fmt.Println("\u2272")
}

