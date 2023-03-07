package ast

import (
	"math"
	"testing"
)

func checkFloat64SlicesEqual(t *testing.T, X, Y []float64) {
	const epsilon float64 = 0.0001
	if len(X) != len(Y) {
		t.Fatalf("Slices have different lengths: %d vs %d", len(X), len(Y))
	}
	for i, a := range X {
		b := Y[i]
		if math.Abs(a-b) > epsilon {
			t.Fatalf("Slices differ at index %d: %f vs %f", i, a, b)
		}
	}
}

func checkFloat32SlicesEqual(t *testing.T, X, Y []float32) {
	const epsilon float32 = 0.0001
	if len(X) != len(Y) {
		t.Fatalf("Slices have different lengths: %d vs %d", len(X), len(Y))
	}
	for i, a := range X {
		b := Y[i]
		if math.Abs(float64(a-b)) > float64(epsilon) {
			t.Fatalf("Slices differ at index %d: %f vs %f", i, a, b)
		}
	}
}

/*** add() ***/

func TestAddWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	add("foo", 1.0)
}

func TestAddFloat64(t *testing.T) {
	a := float64(1.0)
	b := float64(2.0)
	expected := float64(3.0)
	result := add(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAddFloat32(t *testing.T) {
	a := float32(1.0)
	b := float32(2.0)
	expected := float64(3.0)
	result := add(a, b)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestAddFloatMix(t *testing.T) {
	a := float64(1.0)
	b := float32(2.0)
	expected := float64(3.0)
	result := add(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAddFloatMix2(t *testing.T) {
	a := float32(1.0)
	b := float64(2.0)
	expected := float64(3.0)
	result := add(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAddFloat64Vec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{3.0, 5.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloat64VecBroadcastFloat64(t *testing.T) {
	a := float64(1.0)
	b := []float64{2.0, 3.0}
	expected := []float64{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloat64VecBroadcastFloat32(t *testing.T) {
	a := float32(1.0)
	b := []float64{2.0, 3.0}
	expected := []float64{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloat64VecBroadcastFloat64_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float64(1.0)
	expected := []float64{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloat64VecBroadcastFloat32_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float32(1.0)
	expected := []float64{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloat32Vec(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float32{3.0, 5.0}
	result := add(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestAddFloat32BroadcastFloat32(t *testing.T) {
	a := float32(1.0)
	b := []float32{2.0, 3.0}
	expected := []float32{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestAddFloat32BroadcastFloat64(t *testing.T) {
	a := float64(1.0)
	b := []float32{2.0, 3.0}
	expected := []float64{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloat32BroadcastFloat32_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float32(1.0)
	expected := []float32{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestAddFloat32BroadcastFloat64_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float64(1.0)
	expected := []float64{3.0, 4.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloatMixVec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float64{3.0, 5.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAddFloatMixVec2(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{3.0, 5.0}
	result := add(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

/*** subtract() ***/

func TestSubWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	subtract("foo", 1.0)
}

func TestSubFloat64(t *testing.T) {
	a := float64(1.0)
	b := float64(2.0)
	expected := float64(-1.0)
	result := subtract(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSubFloat32(t *testing.T) {
	a := float32(1.0)
	b := float32(2.0)
	expected := float64(-1.0)
	result := subtract(a, b)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestSubFloatMix(t *testing.T) {
	a := float64(1.0)
	b := float32(2.0)
	expected := float64(-1.0)
	result := subtract(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSubFloatMix2(t *testing.T) {
	a := float32(1.0)
	b := float64(2.0)
	expected := float64(-1.0)
	result := subtract(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSubFloat64Vec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{-1.0, -1.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloat64VecBroadcastFloat64(t *testing.T) {
	a := float64(1.0)
	b := []float64{2.0, 3.0}
	expected := []float64{-1.0, -2.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloat64VecBroadcastFloat32(t *testing.T) {
	a := float32(1.0)
	b := []float64{2.0, 3.0}
	expected := []float64{-1.0, -2.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloat64VecBroadcastFloat64_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float64(1.0)
	expected := []float64{1.0, 2.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloat64VecBroadcastFloat32_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float32(1.0)
	expected := []float64{1.0, 2.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloat32Vec(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float32{-1.0, -1.0}
	result := subtract(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestSubFloat32BroadcastFloat32(t *testing.T) {
	a := float32(1.0)
	b := []float32{2.0, 3.0}
	expected := []float32{-1.0, -2.0}
	result := subtract(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestSubFloat32BroadcastFloat64(t *testing.T) {
	a := float64(1.0)
	b := []float32{2.0, 3.0}
	expected := []float64{-1.0, -2.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloat32BroadcastFloat32_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float32(1.0)
	expected := []float32{1.0, 2.0}
	result := subtract(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestSubFloat32BroadcastFloat64_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float64(1.0)
	expected := []float64{1.0, 2.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloatMixVec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float64{-1.0, -1.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSubFloatMixVec2(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{-1.0, -1.0}
	result := subtract(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

/*** multiply() ***/

func TestMulWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	multiply("foo", 1.0)
}

func TestMulFloat64(t *testing.T) {
	a := float64(1.0)
	b := float64(2.0)
	expected := float64(2.0)
	result := multiply(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMulFloat32(t *testing.T) {
	a := float32(1.0)
	b := float32(2.0)
	expected := float64(2.0)
	result := multiply(a, b)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestMulFloatMix(t *testing.T) {
	a := float64(1.0)
	b := float32(2.0)
	expected := float64(2.0)
	result := multiply(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMulFloatMix2(t *testing.T) {
	a := float32(1.0)
	b := float64(2.0)
	expected := float64(2.0)
	result := multiply(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMulFloat64Vec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{2.0, 6.0}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloat64VecBroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloat64VecBroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloat64VecBroadcastFloat64_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloat64VecBroadcastFloat32_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float32(0.5)
	expected := []float64{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloat32Vec(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float32{2.0, 6.0}
	result := multiply(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMulFloat32BroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float32{2.0, 3.0}
	expected := []float32{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMulFloat32BroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float32{2.0, 3.0}
	expected := []float64{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloat32BroadcastFloat32_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float32(0.5)
	expected := []float32{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMulFloat32BroadcastFloat64_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{1.0, 1.5}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloatMixVec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float64{2.0, 6.0}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMulFloatMixVec2(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{2.0, 6.0}
	result := multiply(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

/*** divide() ***/

func TestDivWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	divide("foo", 1.0)
}

func TestDivFloat64(t *testing.T) {
	a := float64(1.0)
	b := float64(2.0)
	expected := float64(0.5)
	result := divide(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestDivFloat32(t *testing.T) {
	a := float32(1.0)
	b := float32(2.0)
	expected := float64(0.5)
	result := divide(a, b)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestDivFloatMix(t *testing.T) {
	a := float64(1.0)
	b := float32(2.0)
	expected := float64(0.5)
	result := divide(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestDivFloatMix2(t *testing.T) {
	a := float32(1.0)
	b := float64(2.0)
	expected := float64(0.5)
	result := divide(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestDivFloat64Vec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{0.5, 0.66666}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloat64VecBroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{0.25, 0.16666}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloat64VecBroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{0.25, 0.16666}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloat64VecBroadcastFloat64_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{4.0, 6.0}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloat64VecBroadcastFloat32_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float32(0.5)
	expected := []float64{4.0, 6.0}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloat32Vec(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float32{0.5, 0.66666}
	result := divide(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestDivFloat32BroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float32{2.0, 3.0}
	expected := []float32{0.25, 0.16666}
	result := divide(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestDivFloat32BroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float32{2.0, 3.0}
	expected := []float64{0.25, 0.16666}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloat32BroadcastFloat32_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float32(0.5)
	expected := []float32{4.0, 6.0}
	result := divide(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestDivFloat32BroadcastFloat64_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{4.0, 6.0}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloatMixVec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float64{0.5, 0.66666}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestDivFloatMixVec2(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{0.5, 0.66666}
	result := divide(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

/*** min() ***/

func TestMinWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	min("foo", 1.0)
}

func TestMinFloat64(t *testing.T) {
	a := float64(1.0)
	b := float64(2.0)
	expected := float64(1.0)
	result := min(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMinFloat32(t *testing.T) {
	a := float32(1.0)
	b := float32(2.0)
	expected := float64(1.0)
	result := min(a, b)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestMinFloatMix(t *testing.T) {
	a := float64(1.0)
	b := float32(2.0)
	expected := float64(1.0)
	result := min(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMinFloatMix2(t *testing.T) {
	a := float32(1.0)
	b := float64(2.0)
	expected := float64(1.0)
	result := min(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMinFloat64Vec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{1.0, 2.0}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloat64VecBroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloat64VecBroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloat64VecBroadcastFloat64_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloat64VecBroadcastFloat32_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float32(0.5)
	expected := []float64{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloat32Vec(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float32{1.0, 2.0}
	result := min(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMinFloat32BroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float32{2.0, 3.0}
	expected := []float32{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMinFloat32BroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float32{2.0, 3.0}
	expected := []float64{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloat32BroadcastFloat32_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float32(0.5)
	expected := []float32{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMinFloat32BroadcastFloat64_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{0.5, 0.5}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloatMixVec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float64{1.0, 2.0}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMinFloatMixVec2(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{1.0, 2.0}
	result := min(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

/*** max() ***/

func TestMaxWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	max("foo", 1.0)
}

func TestMaxFloat64(t *testing.T) {
	a := float64(1.0)
	b := float64(2.0)
	expected := float64(2.0)
	result := max(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMaxFloat32(t *testing.T) {
	a := float32(1.0)
	b := float32(2.0)
	expected := float64(2.0)
	result := max(a, b)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestMaxFloatMix(t *testing.T) {
	a := float64(1.0)
	b := float32(2.0)
	expected := float64(2.0)
	result := max(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMaxFloatMix2(t *testing.T) {
	a := float32(1.0)
	b := float64(2.0)
	expected := float64(2.0)
	result := max(a, b)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestMaxFloat64Vec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloat64VecBroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloat64VecBroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float64{2.0, 3.0}
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloat64VecBroadcastFloat64_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloat64VecBroadcastFloat32_(t *testing.T) {
	a := []float64{2.0, 3.0}
	b := float32(0.5)
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloat32Vec(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float32{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMaxFloat32BroadcastFloat32(t *testing.T) {
	a := float32(0.5)
	b := []float32{2.0, 3.0}
	expected := []float32{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMaxFloat32BroadcastFloat64(t *testing.T) {
	a := float64(0.5)
	b := []float32{2.0, 3.0}
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloat32BroadcastFloat32_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float32(0.5)
	expected := []float32{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestMaxFloat32BroadcastFloat64_(t *testing.T) {
	a := []float32{2.0, 3.0}
	b := float64(0.5)
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloatMixVec(t *testing.T) {
	a := []float64{1.0, 2.0}
	b := []float32{2.0, 3.0}
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestMaxFloatMixVec2(t *testing.T) {
	a := []float32{1.0, 2.0}
	b := []float64{2.0, 3.0}
	expected := []float64{2.0, 3.0}
	result := max(a, b)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

/*** sum() ***/

func TestSumWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	sum("foo")
}

func TestSumFloat64Vec(t *testing.T) {
	a := []float64{-1.0, 1.0}
	expected := 0.0
	result := sum(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	const epsilon float64 = 0.0001
	if math.Abs(result.(float64)-expected) > float64(epsilon) {
		t.Fatalf("output value differs, got: %f expected %f", result.(float64), expected)
	}
}

func TestSumFloat32Vec(t *testing.T) {
	a := []float32{-1.0, 1.0}
	expected := float32(0.0)
	result := sum(a)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	const epsilon float32 = 0.0001
	if math.Abs(float64(result.(float32)-expected)) > float64(epsilon) {
		t.Fatalf("output value differs, got: %f expected %f", result.(float32), expected)
	}
}

/*** abs() ***/

func TestAbsWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	abs("foo")
}

func TestAbsFloat64(t *testing.T) {
	a := float64(-1.0)
	expected := float64(1.0)
	result := abs(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(result.(float64)-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAbsFloat32(t *testing.T) {
	a := float32(-1.0)
	expected := float64(1.0)
	result := abs(a)
	if _, ok := result.(float32); !ok {
		t.Errorf("output value is not float32")
	}
	if math.Abs(float64(result.(float32))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float32))
	}
}

func TestAbsFloat64Vec(t *testing.T) {
	a := []float64{-1.0, 1.0}
	expected := []float64{1.0, 1.0}
	result := abs(a)
	if _, ok := result.([]float64); !ok {
		t.Errorf("output value is not []float64")
	}
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAbsFloat32Vec(t *testing.T) {
	a := []float32{-1.0, 1.0}
	expected := []float32{1.0, 1.0}
	result := abs(a)
	if _, ok := result.([]float32); !ok {
		t.Errorf("output value is not []float32")
	}
	checkFloat32SlicesEqual(t, result.([]float32), expected)
}

func TestAcosFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Acos(a)
	result := acos(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAcosFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Acos(float64(a))
	result := acos(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAcosFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Acos(0.5), math.Acos(0.7)}
	result := acos(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAcosFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Acos(0.5), math.Acos(0.7)}
	result := acos(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAcosBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	acos("ko")
}

func TestAcoshFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Acosh(a)
	result := acosh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAcoshFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Acosh(float64(a))
	result := acosh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAcoshFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Acosh(0.5), math.Acosh(0.7)}
	result := acosh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAcoshFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Acosh(0.5), math.Acosh(0.7)}
	result := acosh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAcoshBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	acosh("ko")
}

func TestAsinFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Asin(a)
	result := asin(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAsinFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Asin(float64(a))
	result := asin(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAsinFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Asin(0.5), math.Asin(0.7)}
	result := asin(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAsinFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Asin(0.5), math.Asin(0.7)}
	result := asin(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAsinBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	asin("ko")
}

func TestAsinhFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Asinh(a)
	result := asinh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAsinhFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Asinh(float64(a))
	result := asinh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAsinhFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Asinh(0.5), math.Asinh(0.7)}
	result := asinh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAsinhFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Asinh(0.5), math.Asinh(0.7)}
	result := asinh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAsinhBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	asinh("ko")
}

func TestAtanFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Atan(a)
	result := atan(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAtanFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Atan(float64(a))
	result := atan(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAtanFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Atan(0.5), math.Atan(0.7)}
	result := atan(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAtanFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Atan(0.5), math.Atan(0.7)}
	result := atan(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAtanBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	atan("ko")
}

func TestAtanhFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Atanh(a)
	result := atanh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAtanhFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Atanh(float64(a))
	result := atanh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestAtanhFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Atanh(0.5), math.Atanh(0.7)}
	result := atanh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAtanhFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Atanh(0.5), math.Atanh(0.7)}
	result := atanh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestAtanhBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	atanh("ko")
}

func TestCbrtFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Cbrt(a)
	result := cbrt(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCbrtFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Cbrt(float64(a))
	result := cbrt(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCbrtFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Cbrt(0.5), math.Cbrt(0.7)}
	result := cbrt(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCbrtFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Cbrt(0.5), math.Cbrt(0.7)}
	result := cbrt(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCbrtBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cbrt("ko")
}

func TestCeilFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Ceil(a)
	result := ceil(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCeilFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Ceil(float64(a))
	result := ceil(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCeilFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Ceil(0.5), math.Ceil(0.7)}
	result := ceil(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCeilFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Ceil(0.5), math.Ceil(0.7)}
	result := ceil(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCeilBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ceil("ko")
}

func TestCosFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Cos(a)
	result := cos(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCosFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Cos(float64(a))
	result := cos(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCosFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Cos(0.5), math.Cos(0.7)}
	result := cos(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCosFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Cos(0.5), math.Cos(0.7)}
	result := cos(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCosBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cos("ko")
}

func TestCoshFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Cosh(a)
	result := cosh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCoshFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Cosh(float64(a))
	result := cosh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestCoshFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Cosh(0.5), math.Cosh(0.7)}
	result := cosh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCoshFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Cosh(0.5), math.Cosh(0.7)}
	result := cosh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestCoshBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cosh("ko")
}

func TestErfFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Erf(a)
	result := erf(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Erf(float64(a))
	result := erf(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Erf(0.5), math.Erf(0.7)}
	result := erf(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Erf(0.5), math.Erf(0.7)}
	result := erf(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	erf("ko")
}

func TestErfcFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Erfc(a)
	result := erfc(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfcFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Erfc(float64(a))
	result := erfc(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfcFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Erfc(0.5), math.Erfc(0.7)}
	result := erfc(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfcFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Erfc(0.5), math.Erfc(0.7)}
	result := erfc(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfcBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	erfc("ko")
}

func TestErfcinvFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Erfcinv(a)
	result := erfcinv(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfcinvFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Erfcinv(float64(a))
	result := erfcinv(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfcinvFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Erfcinv(0.5), math.Erfcinv(0.7)}
	result := erfcinv(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfcinvFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Erfcinv(0.5), math.Erfcinv(0.7)}
	result := erfcinv(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfcinvBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	erfcinv("ko")
}

func TestErfinvFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Erfinv(a)
	result := erfinv(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfinvFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Erfinv(float64(a))
	result := erfinv(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestErfinvFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Erfinv(0.5), math.Erfinv(0.7)}
	result := erfinv(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfinvFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Erfinv(0.5), math.Erfinv(0.7)}
	result := erfinv(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestErfinvBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	erfinv("ko")
}

func TestExpFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Exp(a)
	result := exp(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestExpFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Exp(float64(a))
	result := exp(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestExpFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Exp(0.5), math.Exp(0.7)}
	result := exp(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestExpFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Exp(0.5), math.Exp(0.7)}
	result := exp(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestExpBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	exp("ko")
}

func TestExp2Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.Exp2(a)
	result := exp2(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestExp2Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.Exp2(float64(a))
	result := exp2(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestExp2Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Exp2(0.5), math.Exp2(0.7)}
	result := exp2(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestExp2Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Exp2(0.5), math.Exp2(0.7)}
	result := exp2(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestExp2BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	exp2("ko")
}

func TestExpm1Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.Expm1(a)
	result := expm1(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestExpm1Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.Expm1(float64(a))
	result := expm1(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestExpm1Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Expm1(0.5), math.Expm1(0.7)}
	result := expm1(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestExpm1Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Expm1(0.5), math.Expm1(0.7)}
	result := expm1(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestExpm1BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	expm1("ko")
}

func TestFloorFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Floor(a)
	result := floor(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestFloorFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Floor(float64(a))
	result := floor(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestFloorFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Floor(0.5), math.Floor(0.7)}
	result := floor(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestFloorFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Floor(0.5), math.Floor(0.7)}
	result := floor(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestFloorBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	floor("ko")
}

func TestGammaFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Gamma(a)
	result := gamma(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestGammaFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Gamma(float64(a))
	result := gamma(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestGammaFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Gamma(0.5), math.Gamma(0.7)}
	result := gamma(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestGammaFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Gamma(0.5), math.Gamma(0.7)}
	result := gamma(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestGammaBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	gamma("ko")
}

func TestJ0Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.J0(a)
	result := j0(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestJ0Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.J0(float64(a))
	result := j0(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestJ0Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.J0(0.5), math.J0(0.7)}
	result := j0(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestJ0Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.J0(0.5), math.J0(0.7)}
	result := j0(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestJ0BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	j0("ko")
}

func TestJ1Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.J1(a)
	result := j1(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestJ1Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.J1(float64(a))
	result := j1(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestJ1Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.J1(0.5), math.J1(0.7)}
	result := j1(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestJ1Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.J1(0.5), math.J1(0.7)}
	result := j1(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestJ1BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	j1("ko")
}

func TestLogFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Log(a)
	result := log(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLogFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Log(float64(a))
	result := log(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLogFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Log(0.5), math.Log(0.7)}
	result := log(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLogFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Log(0.5), math.Log(0.7)}
	result := log(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLogBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	log("ko")
}

func TestLog10Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.Log10(a)
	result := log10(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLog10Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.Log10(float64(a))
	result := log10(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLog10Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Log10(0.5), math.Log10(0.7)}
	result := log10(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLog10Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Log10(0.5), math.Log10(0.7)}
	result := log10(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLog10BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	log10("ko")
}

func TestLog1pFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Log1p(a)
	result := log1p(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLog1pFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Log1p(float64(a))
	result := log1p(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLog1pFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Log1p(0.5), math.Log1p(0.7)}
	result := log1p(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLog1pFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Log1p(0.5), math.Log1p(0.7)}
	result := log1p(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLog1pBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	log1p("ko")
}

func TestLog2Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.Log2(a)
	result := log2(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLog2Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.Log2(float64(a))
	result := log2(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLog2Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Log2(0.5), math.Log2(0.7)}
	result := log2(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLog2Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Log2(0.5), math.Log2(0.7)}
	result := log2(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLog2BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	log2("ko")
}

func TestLogbFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Logb(a)
	result := logb(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLogbFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Logb(float64(a))
	result := logb(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestLogbFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Logb(0.5), math.Logb(0.7)}
	result := logb(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLogbFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Logb(0.5), math.Logb(0.7)}
	result := logb(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestLogbBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	logb("ko")
}

func TestRoundFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Round(a)
	result := round(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestRoundFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Round(float64(a))
	result := round(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestRoundFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Round(0.5), math.Round(0.7)}
	result := round(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestRoundFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Round(0.5), math.Round(0.7)}
	result := round(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestRoundBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	round("ko")
}

func TestSinFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Sin(a)
	result := sin(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSinFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Sin(float64(a))
	result := sin(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSinFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Sin(0.5), math.Sin(0.7)}
	result := sin(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSinFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Sin(0.5), math.Sin(0.7)}
	result := sin(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSinBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	sin("ko")
}

func TestSinhFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Sinh(a)
	result := sinh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSinhFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Sinh(float64(a))
	result := sinh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSinhFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Sinh(0.5), math.Sinh(0.7)}
	result := sinh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSinhFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Sinh(0.5), math.Sinh(0.7)}
	result := sinh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSinhBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	sinh("ko")
}

func TestSqrtFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Sqrt(a)
	result := sqrt(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSqrtFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Sqrt(float64(a))
	result := sqrt(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestSqrtFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Sqrt(0.5), math.Sqrt(0.7)}
	result := sqrt(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSqrtFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Sqrt(0.5), math.Sqrt(0.7)}
	result := sqrt(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestSqrtBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	sqrt("ko")
}

func TestTanFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Tan(a)
	result := tan(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestTanFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Tan(float64(a))
	result := tan(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestTanFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Tan(0.5), math.Tan(0.7)}
	result := tan(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestTanFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Tan(0.5), math.Tan(0.7)}
	result := tan(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestTanBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	tan("ko")
}

func TestTanhFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Tanh(a)
	result := tanh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestTanhFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Tanh(float64(a))
	result := tanh(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestTanhFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Tanh(0.5), math.Tanh(0.7)}
	result := tanh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestTanhFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Tanh(0.5), math.Tanh(0.7)}
	result := tanh(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestTanhBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	tanh("ko")
}

func TestTruncFloat64(t *testing.T) {
	a := float64(0.5)
	expected := math.Trunc(a)
	result := trunc(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestTruncFloat32(t *testing.T) {
	a := float32(0.5)
	expected := math.Trunc(float64(a))
	result := trunc(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestTruncFloat32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Trunc(0.5), math.Trunc(0.7)}
	result := trunc(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestTruncFloat64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Trunc(0.5), math.Trunc(0.7)}
	result := trunc(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestTruncBadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	trunc("ko")
}

func TestY0Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.Y0(a)
	result := y0(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestY0Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.Y0(float64(a))
	result := y0(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestY0Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Y0(0.5), math.Y0(0.7)}
	result := y0(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestY0Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Y0(0.5), math.Y0(0.7)}
	result := y0(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestY0BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	y0("ko")
}

func TestY1Float64(t *testing.T) {
	a := float64(0.5)
	expected := math.Y1(a)
	result := y1(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestY1Float32(t *testing.T) {
	a := float32(0.5)
	expected := math.Y1(float64(a))
	result := y1(a)
	if _, ok := result.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	if math.Abs(float64(result.(float64))-expected) > 1e-6 {
		t.Errorf("expected %f but got %f", expected, result.(float64))
	}
}

func TestY1Float32Vec(t *testing.T) {
	a := []float32{0.5, 0.7}
	expected := []float64{math.Y1(0.5), math.Y1(0.7)}
	result := y1(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestY1Float64Vec(t *testing.T) {
	a := []float64{0.5, 0.7}
	expected := []float64{math.Y1(0.5), math.Y1(0.7)}
	result := y1(a)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestY1BadArg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	y1("ko")
}
