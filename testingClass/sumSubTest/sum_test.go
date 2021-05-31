package sum

import "testing"

// go test -v // to run tests
func TestSumInt(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int // field names
		want    int
	}{
		{
			name:    "one to five",
			numbers: []int{1, 2, 3, 4, 5},
			want:    15,
		}, // test cases //tc
		{name: "no numbers", numbers: nil, want: 0},
		{name: "one and minus one", numbers: []int{1, -1}, want: 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			got := SumInt(tc.numbers)

			if got != tc.want {
				t.Fatalf("sum of %v should be %v; got %v", tc.numbers, tc.want, got)
			}
		})
	}

}

// go test -run SumInt/one // it's pattern matching
