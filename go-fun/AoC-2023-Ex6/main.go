package main

import (
  "fmt"
  "bufio"
)

type SortBy []Type
  
func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func Benchmark(b *testing.B) {
  for i := 0; i < b.N; i++ {
      
  }
}
func main(){}
