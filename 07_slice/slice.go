package main

import "fmt"

func main() {
  var numbers1 []int /* unspecified size */
  var numbers2 []int = make([]int, 5, 5) /* a slice of length 5 and capacity 5 */
  var numbers3 []int
  var numbers4 []int

  numbers1 = append(numbers1, 6)
  numbers2[2] = 4

  numbers3 = make([]int, len(numbers1), cap(numbers1))
  copy(numbers3, numbers1) /* copy content of numbers1 to numbers3 */

  numbers4 = numbers2[:3] // range 0-3

  printSlice("numbers1", numbers1)
  printSlice("numbers2", numbers2)
  printSlice("numbers3", numbers3)
  printSlice("numbers4", numbers4)

}

func printSlice(name string, x []int) {
  fmt.Printf("%s :: len=%d cap=%d slice=%v\n", name, len(x), cap(x), x)
}
