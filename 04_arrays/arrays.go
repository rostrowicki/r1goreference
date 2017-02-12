package main

import (
  "fmt"
  "strconv"
)

func main() {

  fmt.Printf("Arrays...\n")

// -------------- integers
  var drawer = [5] float32 {1.2,2.3,3.4,5.6,7.8}
  for index, value := range drawer {
    fmt.Printf("value of element %d is %f \n", index, value)
  }

// -------------- strings
  var words = [5] string {"A", "B", "C", "D", "E"}
  for index,value := range words {
    fmt.Printf("Word under %d is %s\n", index, value)
  }

// -------------- passing to function
fmt.Printf("Average number is: %f\n", average(drawer))

// -------------- sllice passing to function
var wardrobe = []uint8 {1,2,3,4,5}
averageSlice(wardrobe)
}

// using Array as argument
func average( arr [5]float32) (float32) {
  var sum float32
  var size int
  for index,value := range arr {
    sum+=value
    size = index + 1
  }
  fmt.Println("Value of size is " + strconv.Itoa(size)) // converts to string
  return sum / float32(size)
}

// using Slice as argument
func averageSlice(arr []uint8 )(uint8) {
  var sum, size uint8

  for index,value := range arr {
    sum += value
    size = uint8(index + 1)
  }
  fmt.Println("Returned value is " + strconv.Itoa(int(sum/size))) // uint8 > int > string
  return sum / size
}
