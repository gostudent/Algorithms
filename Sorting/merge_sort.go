package main

import "fmt"

func Merge(a1 []int, a2 []int) []int {

  var r = make([]int, len(a1) + len(a2))
  var i = 0
  var j = 0

  for i < len(a1) && j < len(a2) {
  
    if a1[i] <= a2[j] {
      r[i+j] = a1[i]
      i++
    } else {
      r[i+j] = a2[j]
      j++
    }
    
  }

  for i < len(a1) { r[i+j] = a1[i]; i++ }
  for j < len(a2) { r[i+j] = a2[j]; j++ }

  return r
  
}

func Merge_Sort(items []int) []int {

  if len(items) < 2 {
    return items
    
  }

  var middle = len(items) / 2
  var a = Merge_Sort(items[:middle])
  var b = Merge_Sort(items[middle:])
  return Merge(a, b)
  
}

func main () {

  fmt.Print(Merge_Sort([]int{ 10, 92, 83, 43,45, 64, 744, 33, 22, 14,2344,56743 }), "\n")
  
}