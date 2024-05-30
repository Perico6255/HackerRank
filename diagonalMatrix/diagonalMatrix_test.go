package diagonalmatrix

import "testing"

func TestDiagonalMatrix(t *testing.T) {
  arr := [][]int32{
    {11,2,4},
    {4,5,6},
    {10,8,-12},
  }
  output := DiagonalDifference(arr)

  if 15!=output {
    t.Errorf("DiagonalDifference Fatal \n EXPECTED: %v \n OUTPUT %v", 15,output)
    
  }
  
}
