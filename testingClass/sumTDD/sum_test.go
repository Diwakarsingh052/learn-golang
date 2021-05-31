package sum

import "testing"

// go test -v // to run tests
func TestSumInt(t *testing.T) {
	tt := []struct {
		numbers []int // field names
		want    int
	}{
		{numbers: []int{1, 2, 3, 4, 5}, want: 15}, // test cases //tc
		{numbers: nil, want: 0},
		{numbers: []int{1, -1}, want: 0},
	}

	for _, tc := range tt {

		got := SumInt(tc.numbers)

		if got != tc.want {
			t.Errorf("sum of %v want %v; got %v", tc.numbers, tc.want, got)
		}

	}

}
func TestFoo(t *testing.T) {
	t.Errorf("Foo always fail")
}

// go test -run SumInt
