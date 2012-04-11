package vector

import "testing"

func TestDot(t *testing.T) {
	// A list of our dot product testcases.
	cases := [...]struct {
		One, Two Vec3
		Output   float64
	}{
		{Vec3{0, 0, 0}, Vec3{0, 0, 0}, 0},
		{Vec3{1, 1, 1}, Vec3{0, 0, 0}, 0},
		{Vec3{0, 0, 0}, Vec3{1, 1, 1}, 0},
		{Vec3{1, 1, 1}, Vec3{1, 1, 1}, 3},
		{Vec3{10, 1, 1}, Vec3{1, 1, 1}, 12},
		{Vec3{0, 10, 0}, Vec3{0, 10, 0}, 100},
		{Vec3{3, 4, 5}, Vec3{1, 1, 1}, 12},
		{Vec3{2, 2, 3}, Vec3{3, 3, 2}, 18},
		{Vec3{1, 1, 1}, Vec3{-1, -1, -1}, -3},
	}

	for _, test := range cases {
		if test.Output != Dot(&test.One, &test.Two) {
			t.Error("Expected dot product of", test.One, "and",
				test.Two, "to be", test.Output)
		}
	}
}
