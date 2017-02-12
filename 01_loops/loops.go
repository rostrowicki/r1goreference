package main

import "fmt"

func main() {
    fmt.Printf("Loops..\n")

    // -------------------------
    var x int
    for x =1; x<10; x++ {
        fmt.Printf("*")
    }
    fmt.Printf("\n")

    // -------------------------
    var y int = 20
    for x<y {
      fmt.Printf("+")
      x++
    }
    fmt.Printf("\n")

    // -------------------------
    numbers := [4]int{7,8,9,6}
    for index,value:= range numbers {
      fmt.Printf("value of x = %d at %d \n", value, index)
    }
    fmt.Printf("\n")

    fmt.Printf("TERMINATED\n")
}
