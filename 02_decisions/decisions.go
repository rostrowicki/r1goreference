package main

import "fmt"

func main() {
  fmt.Printf("Decision %s\n", "makers")

  var x int = 100
  var y int = 200

if (x == 100) {
  if ( y == 200) {
    fmt.Printf("\nValue of x is 100 and y is 200\n")
  }
}


// -----------------------
var z int = 3

switch z {
case 1:
   fmt.Printf("Z is ONE\n")
 case 2:
  fmt.Printf("Z is TWO\n")
default:
  fmt.Printf("Z is %d\n", z)
}

}
