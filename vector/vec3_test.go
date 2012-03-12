package vector

import "testing"

func TestLength(t *testing.T) {
	v1 := &Vec3{3.0, 4.0, 0.0}
	vLen := v1.Length()
	if vLen != 5.0 {
		t.Errorf("Expected %f, got:", 5.0, vLen)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := &Vec3{3.0, 4.0, 5.0}
	v2 := &Vec3{6.0, 7.0, 8.0}

	expected := 86.0
	value := v1.DotProduct(v2)

	if expected != value {
		t.Errorf("Expected %f, got:", expected, value)
	}

	expected = 50.0
	value = v1.DotProduct(v1)

	if expected != value {
		t.Errorf("Expected %f, got:", expected, value)
	}
}
