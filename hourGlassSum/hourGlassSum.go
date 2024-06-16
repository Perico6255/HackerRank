package hourglasssum

import "fmt"

func HourGlassSum(arr [][]int32)int32{
  
  var max  int32
  var now  int32
  for j := 0; j < 4; j++ {
    for i := 0; i < 4; i++ {
      now = arr[j][i] + arr[j][i+1] + arr[j][i+2] + arr[j+1][i+1] + arr[j+2][i] + arr[j+2][i+1] + arr[j+2][i+2]
      fmt.Println(now)
      if max < now || (i==0 && j==0)  {
        max = now
      }
    }
    
  }

  return max
}

