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

func TestEvaluateAcos(t *testing.T) {
	ast, _ := ParseExpr(`acos(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Acos(-1.0), math.Acos(3.0), math.Acos(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateAcosh(t *testing.T) {
	ast, _ := ParseExpr(`acosh(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Acosh(-1.0), math.Acosh(3.0), math.Acosh(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateAsin(t *testing.T) {
	ast, _ := ParseExpr(`asin(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Asin(-1.0), math.Asin(3.0), math.Asin(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateAsinh(t *testing.T) {
	ast, _ := ParseExpr(`asinh(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Asinh(-1.0), math.Asinh(3.0), math.Asinh(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateAtan(t *testing.T) {
	ast, _ := ParseExpr(`atan(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Atan(-1.0), math.Atan(3.0), math.Atan(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateAtanh(t *testing.T) {
	ast, _ := ParseExpr(`atanh(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Atanh(-1.0), math.Atanh(3.0), math.Atanh(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateCbrt(t *testing.T) {
	ast, _ := ParseExpr(`cbrt(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Cbrt(-1.0), math.Cbrt(3.0), math.Cbrt(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateCeil(t *testing.T) {
	ast, _ := ParseExpr(`ceil(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Ceil(-1.0), math.Ceil(3.0), math.Ceil(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateCos(t *testing.T) {
	ast, _ := ParseExpr(`cos(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Cos(-1.0), math.Cos(3.0), math.Cos(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateCosh(t *testing.T) {
	ast, _ := ParseExpr(`cosh(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Cosh(-1.0), math.Cosh(3.0), math.Cosh(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateErf(t *testing.T) {
	ast, _ := ParseExpr(`erf(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Erf(-1.0), math.Erf(3.0), math.Erf(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateErfc(t *testing.T) {
	ast, _ := ParseExpr(`erfc(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Erfc(-1.0), math.Erfc(3.0), math.Erfc(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateErfcinv(t *testing.T) {
	ast, _ := ParseExpr(`erfcinv(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Erfcinv(-1.0), math.Erfcinv(3.0), math.Erfcinv(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateErfinv(t *testing.T) {
	ast, _ := ParseExpr(`erfinv(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Erfinv(-1.0), math.Erfinv(3.0), math.Erfinv(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateExp(t *testing.T) {
	ast, _ := ParseExpr(`exp(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Exp(-1.0), math.Exp(3.0), math.Exp(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateExp2(t *testing.T) {
	ast, _ := ParseExpr(`exp2(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Exp2(-1.0), math.Exp2(3.0), math.Exp2(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateExpm1(t *testing.T) {
	ast, _ := ParseExpr(`expm1(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Expm1(-1.0), math.Expm1(3.0), math.Expm1(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateFloor(t *testing.T) {
	ast, _ := ParseExpr(`floor(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Floor(-1.0), math.Floor(3.0), math.Floor(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateGamma(t *testing.T) {
	ast, _ := ParseExpr(`gamma(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Gamma(-1.0), math.Gamma(3.0), math.Gamma(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateJ0(t *testing.T) {
	ast, _ := ParseExpr(`j0(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.J0(-1.0), math.J0(3.0), math.J0(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateJ1(t *testing.T) {
	ast, _ := ParseExpr(`j1(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.J1(-1.0), math.J1(3.0), math.J1(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateLog(t *testing.T) {
	ast, _ := ParseExpr(`log(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Log(-1.0), math.Log(3.0), math.Log(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateLog10(t *testing.T) {
	ast, _ := ParseExpr(`log10(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Log10(-1.0), math.Log10(3.0), math.Log10(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateLog1p(t *testing.T) {
	ast, _ := ParseExpr(`log1p(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Log1p(-1.0), math.Log1p(3.0), math.Log1p(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateLog2(t *testing.T) {
	ast, _ := ParseExpr(`log2(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Log2(-1.0), math.Log2(3.0), math.Log2(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateLogb(t *testing.T) {
	ast, _ := ParseExpr(`logb(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Logb(-1.0), math.Logb(3.0), math.Logb(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateRound(t *testing.T) {
	ast, _ := ParseExpr(`round(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Round(-1.0), math.Round(3.0), math.Round(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateSin(t *testing.T) {
	ast, _ := ParseExpr(`sin(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Sin(-1.0), math.Sin(3.0), math.Sin(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateSinh(t *testing.T) {
	ast, _ := ParseExpr(`sinh(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Sinh(-1.0), math.Sinh(3.0), math.Sinh(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateSqrt(t *testing.T) {
	ast, _ := ParseExpr(`sqrt(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Sqrt(-1.0), math.Sqrt(3.0), math.Sqrt(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateTan(t *testing.T) {
	ast, _ := ParseExpr(`tan(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Tan(-1.0), math.Tan(3.0), math.Tan(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateTanh(t *testing.T) {
	ast, _ := ParseExpr(`tanh(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Tanh(-1.0), math.Tanh(3.0), math.Tanh(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateTrunc(t *testing.T) {
	ast, _ := ParseExpr(`trunc(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Trunc(-1.0), math.Trunc(3.0), math.Trunc(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateY0(t *testing.T) {
	ast, _ := ParseExpr(`y0(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Y0(-1.0), math.Y0(3.0), math.Y0(-2.0)}
	result := Evaluate(ast, vars)
	checkFloat64SlicesEqual(t, result.([]float64), expected)
}

func TestEvaluateY1(t *testing.T) {
	ast, _ := ParseExpr(`y1(X)`)
	vars := &Env{
		"X": []float64{-1.0, 3.0, -2.0},
	}
	expected := []float64{math.Y1(-1.0), math.Y1(3.0), math.Y1(-2.0)}
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
