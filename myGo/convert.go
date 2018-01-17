package main

import "fmt"

func main() {

  var num int
  fmt.Print("Convert a number to binary, hex, and UTF8: ")
  fmt.Scan(&num)
  fmt.Printf("%b \t %x \t %q", num, num, num)
}
