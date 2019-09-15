package main

import "testing"

func TestCreateVectors(t *testing.T) {
	v := Vector{1, 2, 3}

	if v.x != 1 {
		t.Errorf("Expected x = 1, got %f", v.x)
	}

	if v.y != 2 {
		t.Errorf("Expected y = 1, got %f", v.y)
	}

	if v.z != 3 {
		t.Errorf("Expected z = 1, got %f", v.z)
	}
}

func TestAddVectors(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 1, 1}

	got := v1.add(v2)
	expected := Vector{2, 3, 4}

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestAddVectorsEquality(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{1, 1, 1}

	first := v1.add(v2)
	second := v2.add(v1)

	if first != second {
		t.Errorf("Expected addition to be commutative: %v != %v", first, second)
	}
}

func TestScaleVector(t *testing.T) {
	v := Vector{1, 2, 3}
  const scalar float64 = 10

	expected := Vector{10, 20, 30}
	got := v.scale(scalar)

	if expected != got {
		t.Errorf("Incorrect scaling. Expected: %v, Got: %v", expected, got)
	}
}

func TestSubtractVectors(t *testing.T) {
  v1 := Vector{1, 2, 3}
  v2 := Vector{3, 2, 1}

  expected := Vector{-2, 0, 2}
  got := v1.subtract(v2)

  if expected != got {
    t.Errorf("Incorrect subtraction. Expected %v, Got: %v", expected, got)
  }
}
