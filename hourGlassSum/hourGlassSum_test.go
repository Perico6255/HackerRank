package hourglasssum

import (
	tcolors "hackerrank/tColors"
	"reflect"
	"testing"
)


func TestHourGlassSum(t *testing.T) {
  var simulations = []struct{
    input [][]int32
    espected int32
  }{
      {
      [][]int32{
        {1,1,1,0,0,0},
        {0,1,0,0,0,0},
        {1,1,1,0,0,0},
        {0,0,2,4,4,0},
        {0,0,0,2,0,0},
        {0,0,1,2,4,0},
      },19}}

  for _,s :=range simulations{
    output := HourGlassSum(s.input)
    if !reflect.DeepEqual(s.espected, output) {
      t.Errorf(tcolors.Yellow+  "\n FAIL \n"  +"EXPECTED %v \n" + tcolors.Red+ "OUTPUT %v" + tcolors.Reset,s.espected,output)
    }
  }
  
  
}
