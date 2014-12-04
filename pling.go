package main

import (
  "strconv"
  "fmt"
  )

func main() {
  var x int = 28
  var y int = 1755
  var z int = 34

  fmt.Println(raindrops(x))
  fmt.Println(raindrops(y))
  fmt.Println(raindrops(z))
}

func raindrops(num int) string {
  var output string
  if pling(num) { output += "Pling" }
  if plang(num) { output += "Plang" }
  if plong(num) { output += "Plong" }
  if !pling(num) && !plang(num) && !plong(num) { output = strconv.Itoa(num) }
  return output
}

func pling(num int) bool { return num % 3 == 0 }
func plang(num int) bool { return num % 5 == 0 }
func plong(num int) bool { return num % 7 == 0 }
