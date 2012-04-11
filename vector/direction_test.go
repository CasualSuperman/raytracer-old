package vector

import "testing"
//import "math"

func TestScale(t *testing.T) {
	testCases := []struct{
		InVec Direction
		Factor float64
		OutVec Direction
	}{
		{Direction{ 0,  0,  0},   0, Direction{  0,  0,  0}},
		{Direction{ 1,  1,  1},   0, Direction{  0,  0,  0}},
		{Direction{ 0,  0,  0},   1, Direction{  0,  0,  0}},
		{Direction{-1,  1,  1},   0, Direction{  0,  0,  0}},
		{Direction{ 1,  1,  1},   2, Direction{  2,  2,  2}},
		{Direction{ 1,  0,  0},  -1, Direction{ -1,  0,  0}},
		{Direction{ 4,  6,  8}, 0.5, Direction{  2,  3,  4}},
		{Direction{ 2,  3,  4},  -2, Direction{ -4, -6, -8}},
		{Direction{ 2, -3,  0},  -2, Direction{ -4,  6,  0}},
	}

	for _, test := range testCases {
		newVec := test.InVec.Copy()
		newVec.Scale(test.Factor)
		if ! (newVec.X == test.OutVec.X &&
			  newVec.Y == test.OutVec.Y &&
			  newVec.Z == test.OutVec.Z) {
			t.Error("Direction.Scale", test.InVec, "*",
					test.Factor, "!=", test.OutVec)
		}
	}
}
