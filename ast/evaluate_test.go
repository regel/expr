package ast

import (
	"math"
	"os"
	"testing"
)

func TestEvaluate(t *testing.T) {
	vars := &Env{
		"aa": 5.0,
		"bb": []float64{1.0, 5.0, 2.0},
	}
	tests := []struct {
		expr     string
		expected float64
	}{
		{"2 + 3", 5},
		{"2 - 3", -1},
		{"2 * 3", 6},
		{"2 * aa", 10},
		{"2 * bb[1]", 10},
		{"8 / 4", 2},
		{"2 + 3 * 4", 14},
		{"2 * 3 + 4 * 5", 26},
		{"2 * (3 + 4.0) * 5", 70},
		{"2.21 * (3.07 + 4) * 5.001", 78.1391247},
	}

	for _, test := range tests {
		ast, _ := ParseExpr(test.expr)
		PrettyPrint(os.Stdout, ast, " ")
		result := Evaluate(ast, vars)
		if _, ok := result.(float64); !ok {
			t.Errorf("output value is not float64")
		}
		if math.Abs(result.(float64)-test.expected) > 1e-9 {
			t.Errorf("For expression %s, expected %f but got %f", test.expr, test.expected, result.(float64))
		}
	}
}

func TestEvaluateVectorOps(t *testing.T) {
	mixed_vars := []*Env{
		&Env{
			"aa": []float64{-1.0, -3.0, -2.0},
			"bb": []float64{1.0, 5.0, 2.0},
			"c":  10.0,
		},
		&Env{
			"aa": []float64{-1.0, -3.0, -2.0},
			"bb": []float64{1.0, 5.0, 2.0},
			"c":  float32(10.0),
		},
		&Env{
			"aa": []float64{-1.0, -3.0, -2.0},
			"bb": []float32{1.0, 5.0, 2.0},
			"c":  float64(10.0),
		},
		&Env{
			"aa": []float32{-1.0, -3.0, -2.0},
			"bb": []float64{1.0, 5.0, 2.0},
			"c":  float64(10.0),
		},
	}

	tests := []struct {
		expr     string
		expected []float64
	}{
		{"aa + bb", []float64{0.0, 2.0, 0.0}},
		{"bb + aa", []float64{0.0, 2.0, 0.0}},
		{"aa - bb", []float64{-2.0, -8.0, -4.0}},
		{"aa * bb", []float64{-1.0, -15.0, -4.0}},
		{"bb * aa", []float64{-1.0, -15.0, -4.0}},
		{"aa / bb", []float64{-1.0, -0.6, -1.0}},
		{"aa + c", []float64{9.0, 7.0, 8.0}},
		{"c + aa", []float64{9.0, 7.0, 8.0}},
		{"aa - c", []float64{-11.0, -13.0, -12.0}},
		{"aa * c", []float64{-10.0, -30.0, -20.0}},
		{"c * aa", []float64{-10.0, -30.0, -20.0}},
		{"aa / c", []float64{-0.1, -0.3, -0.2}},
	}

	for _, test := range tests {
		ast, _ := ParseExpr(test.expr)
		for _, vars := range mixed_vars {
			result := Evaluate(ast, vars)
			checkFloat64SlicesEqual(t, result.([]float64), test.expected)
		}
	}
}

func TestEvaluateAbs(t *testing.T) {
	ast, _ := ParseExpr(`abs(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{1.0, 3.0, 2.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateAdd(t *testing.T) {
	ast, _ := ParseExpr(`add(X, Y)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
		"Y": []float64{0.0, -1.0, 1.0},
	}
	expected := []float64{-1.0, 2.0, -1.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateSub(t *testing.T) {
	ast, _ := ParseExpr(`sub(X, Y)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
		"Y": []float64{0.0, -1.0, 1.0},
	}
	expected := []float64{-1.0, 4.0, -3.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateMul(t *testing.T) {
	ast, _ := ParseExpr(`mul(X, Y)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
		"Y": []float64{0.0, -1.0, 1.0},
	}
	expected := []float64{0.0, -3.0, -2.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateDiv(t *testing.T) {
	ast, _ := ParseExpr(`div(X, Y)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
		"Y": []float64{1.0, -2.0, 2.0},
	}
	expected := []float64{-1.0, -1.5, -1.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateMin(t *testing.T) {
	ast, _ := ParseExpr(`min(X, Y)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
		"Y": []float64{0.0, -1.0, 1.0},
	}
	expected := []float64{-1.0, -1.0, -2.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateMax(t *testing.T) {
	ast, _ := ParseExpr(`max(X, Y)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
		"Y": []float64{0.0, -1.0, 1.0},
	}
	expected := []float64{0.0, 3.0, 1.0}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateNoEnv(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ast, _ := ParseExpr("aa * 3")
	Evaluate(ast, nil)
}

func TestEvaluateNoKeyInEnv(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	vars := &Env{
		"aa": 1.0,
	}
	ast, _ := ParseExpr("bb * 3")
	Evaluate(ast, vars)
}

func TestEvaluateStaticTypes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	vars := &Env{
		"aa": 1.0,
		"bb": "foo",
	}
	ast, _ := ParseExpr("bb * aa")
	Evaluate(ast, vars)
}

func TestEvaluateNilValueInEnv(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	vars := &Env{
		"aa": nil,
	}
	ast, _ := ParseExpr("aa * 3")
	Evaluate(ast, vars)
}

func TestParseVars(t *testing.T) {
	ast, err := ParseExpr("aa[1] * (bb - dd) / c")
	if err != nil {
		t.Error("got error")
	}
	PrettyPrint(os.Stdout, ast, " ")
}

func TestParseExprErr(t *testing.T) {
	tests := []struct {
		expr     string
		expected error
	}{
		{"(1", &ParseError{at: 1, message: "unbalanced parenthesis"}},
		{"1)", &ParseError{at: 1, message: "unbalanced parenthesis"}},
		{"1* (2", &ParseError{at: 4, message: "unbalanced parenthesis"}},
		{"1 * (2 + 3", &ParseError{at: 5, message: "unbalanced parenthesis"}},
		{"1 * (2 + 3))", &ParseError{at: 11, message: "unbalanced parenthesis"}},
	}
	for _, test := range tests {
		ast, err := ParseExpr(test.expr)
		if err == nil {
			t.Error("Expected error but got nil")
		}
		if err.Error() != test.expected.Error() {
			t.Errorf("Expected error '%s' but got '%s'", test.expected, err)
		}
		if ast != nil {
			t.Errorf("Expected ast to be nil but got %v", ast)
		}
	}
}
