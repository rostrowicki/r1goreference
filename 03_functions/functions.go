package main

import "fmt"

func main() {
  fmt.Printf("Functions...\n")
  funcA(5,7)

  a,b := swap("ONE", "TWO")
  fmt.Println(a,b) // output: TWO ONE
}

func funcA(numX, numY int) int {
  fmt.Printf("X is %d, Y is %d\n", numX, numY)
  return 0
}

func swap (a, b string) (string, string) {
  return b,a
}
