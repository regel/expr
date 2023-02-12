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
