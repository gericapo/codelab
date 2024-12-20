package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

    file, err := os.Open("input.txt")
    if err != nil {
		  fmt.Println("Error opening file:", err)
		  return
	  }
	  defer file.Close()
    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        line := strings.Fields(scanner.Text())
        //fmt.Println(line)
        if len(line) == 0 {
            continue
        }
        nums := make([]int, len(line))
        for i, v := range line {
            n, err := strconv.Atoi(v)
            if err != nil {
                continue
            }
            nums[i] = n
            //fmt.Printf("%d", nums[i])
        }
        if len(nums) < 2 {
            continue
        }
        diffs := make([]int, len(nums)-1)
        safe := true
        for i := 0; i < len(nums)-1; i++ {
            d := nums[i+1] - nums[i]
            if d == 0 || d < -3 || d > 3 {
                safe = false
                break
            }
            diffs[i] = d
        }
        if safe {
            increasing := true
            decreasing := true
            for _, d := range diffs {
                if d <= 0 {
                    increasing = false
                }
                if d >= 0 {
                    decreasing = false
                }
            }
            if (increasing || decreasing) && safe {
                count++
            }
        }
    }
    fmt.Println(count)
}

