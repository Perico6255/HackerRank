package reversearray

import (
	tcolors "hackerrank/tColors"
	"reflect"
	"testing"
)


func TestReverseArray(t *testing.T) {
  var simulations = []struct{
    input []int32
    espected []int32
  }{
      {[]int32{1,2,3},[]int32{3,2,1}},
      {[]int32{1,2,3,4},[]int32{4,3,2,1}},
      {[]int32{3,2,3},[]int32{3,2,3}},
    }
  for _,s :=range simulations{
    output := ReverseArray(s.input)
    if !reflect.DeepEqual(s.espected, output) {
      t.Errorf(tcolors.Yellow+  "\n FAIL \n"  +"EXPECTED %v \n" + tcolors.Red+ "OUTPUT %v" + tcolors.Reset,s.espected,output)
    }
  }
  
  
}
