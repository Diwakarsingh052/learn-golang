package sum

import "testing"

// go test -v // to run tests
func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	s := SumInt(x)

	if s != 15 {
		t.Fatalf("sum of 1 to 5 should be 15; got %v", s)
	}
	s = SumInt([]int{})

	if s!=0 {
		t.Errorf("sum of no numbers should be 0; got %v", s)
	}


}
