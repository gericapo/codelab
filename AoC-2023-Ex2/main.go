package main

import(
  "fmt"
  "strconv"
  "strings"
  "bufio"
  "os"
)

func main(){
  mred, mgreen, mblue := 12, 13, 14
  
  file, err := os.Open("input.txt")
  if err != nil{
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  sumOfValidGameIDs := 0

  for scanner.Scan(){
    line := scanner.Text()
     
    parts:= strings.SplitN(line, ": ", 2)
    gameID, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "Game"))
    cubeSets := strings.Split(parts[1], "; ")

    valid := true
    for _, set:= range cubeSets{
      cubes:= strings.Split(set, ", ")
      red, green, blue := 0,0,0
      
      for _, cube:= range cubes {
        parts := strings.Fields(cube)
        count, _ := strconv.Atoi(parts[0])
        color := parts[1]
        
        switch color {
          case "red":
            red += count
          case "green":
            green += count
          case "blue":
            blue += count

        }
      }
      if red > mred || green > mgreen || blue > mblue {
        valid = false
        break
      }
    }
    if valid {
      sumOfValidGameIDs += gameID
    }
  }
  if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

  fmt.Println("Sum of valid game IDs:", sumOfValidGameIDs)
}
