package main

import "fmt"

func main() {
  fmt.Printf("Pointers...\n");

  var a int = 37
  var x float32 = 3.789123
  var n = "Robert"

  var ip *int           /* pointer to an int */
  var fp *float32       /* pointer to a float */
  var sp *string        /* pointer to a string */

  ip = &a               /* store address of a in pointer variable */
  fp = &x
  sp = &n

  fmt.Printf("Address of variable a: %x\n", &a)
  fmt.Printf("Value of variable a: %d\n", *ip)

  fmt.Printf("Adress of variable x: %x\n", &x)
  fmt.Printf("Value of variable x: %f\n", *fp)

  fmt.Printf("Address of variable n: %x\n", &n)
  fmt.Printf("Value of variable n: %s\n", *sp)

  fmt.Printf("\nAnd now we are changing variables using pointers...\n")
  changeVars(ip, fp, sp);

  fmt.Printf("a = %d\n", a)
  fmt.Printf("x = %f\n", x)
  fmt.Printf("n = %s\n", n)

}


func changeVars(v1 *int, v2 *float32, v3 *string) {
  *v1 = 38
  *v2 = 4.1234567
  *v3 = "Bob"
}
