package expr_test

import (
	"github.com/regel/expr"
	"github.com/regel/expr/ast"
	"github.com/stretchr/testify/require"
	"math"
	"math/rand"
	"testing"

	"time"
)

func generateRandomFloats(n int) []float64 {
	randomFloats := make([]float64, n)
	for i := 0; i < n; i++ {
		randomFloats[i] = rand.Float64()
	}
	return randomFloats
}

func TestCompile(t *testing.T) {
	input := `((((4 * 2) + (3 * 5)) / (8 - 5)) - (7 + 3)) * (2 / (1 + 1))`
	ast, err := expr.Compile(input)
	require.NotNil(t, ast, "expr.Compile returned a nil value")
	require.NoError(t, err, "expr.Compile returned an error")
}

func TestRun(t *testing.T) {
	input := `((((4 * 2) + (3 * 5)) / (8 - 5)) - (7 + 3)) * (2 / (1 + 1))`
	ast, err := expr.Compile(input)
	require.NotNil(t, ast, "expr.Compile returned a nil value")
	require.NoError(t, err, "expr.Compile returned an error")

	expected := -2.333333
	for i := 0; i < 2; i++ {
		out, err := expr.Run(ast, nil)
		require.NotNil(t, out, "expr.Run returned a nil value")
		require.NoError(t, err, "expr.Run returned an error")
		if math.Abs(out.(float64)-expected) > 1e-5 {
			t.Errorf("For expression %s, expected %f but got %f", input, expected, out)
		}
	}
}

func TestEvaluate(t *testing.T) {
	input := `((((4 * 2) + (3 * 5)) / (8 - 5)) - (7 + 3)) * (2 / (1 + 1))`
	out, _ := expr.Evaluate(input, nil)
	if _, ok := out.(float64); !ok {
		t.Errorf("output value is not float64")
	}
	expected := -2.333333
	if math.Abs(out.(float64)-expected) > 1e-5 {
		t.Errorf("For expression %s, expected %f but got %f", input, expected, out)
	}
}

func TestEvaluateRecover(t *testing.T) {
	_, err := expr.Evaluate("aa * 3", nil)
	expected := &expr.EvaluateError{Message: "Cannot evaluate expression. Key 'aa' not found in environment"}
	if err == nil {
		t.Error("Expected error but got nil")
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error '%s' but got '%s'", expected, err)
	}
}

func TestEvaluateMissingOperand(t *testing.T) {
	_, err := expr.Evaluate("aa * ", nil)
	expected := &expr.EvaluateError{Message: "runtime error: index out of range [-1]"}
	if err == nil {
		t.Error("Expected error but got nil")
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error '%s' but got '%s'", expected, err)
	}
}

func TestEvaluateCos(t *testing.T) {
	code := `2 * cos(Features)`
	program, err := expr.Compile(code)
	require.NoError(t, err)
	env := &ast.Env{
		"Features": []float64{-1.0, 0.0, 1.0},
	}
	_, err = expr.Run(program, env)
	require.NoError(t, err)
}

func TestEvaluateSum(t *testing.T) {
	code := `1 - sum(Features) / 2.0`
	program, err := expr.Compile(code)
	require.NoError(t, err)
	env := &ast.Env{
		"Features": []float64{-1.0, 1.0, 1.0},
	}
	actual, err := expr.Run(program, env)
	expected := 0.5
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func BenchmarkFloat_expr(b *testing.B) {
	params := &ast.Env{
		"aa": 1.0,
		"bb": 2.0,
	}
	program, err := expr.Compile(`aa * bb`)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if _, ok := out.(float64); !ok {
		b.Fail()
	}
}

func benchmarkFloat64_expr(b *testing.B, input string, num int) {
	rand.Seed(time.Now().UnixNano())
	params := &ast.Env{
		"aa": generateRandomFloats(num),
		"bb": generateRandomFloats(num),
	}
	program, err := expr.Compile(input)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if _, ok := out.([]float64); !ok {
		b.Fail()
	}
	if len(out.([]float64)) != num {
		b.Fail()
	}
}

func BenchmarkFloatVecMul_expr(b *testing.B) {
	benchmarkFloat64_expr(b, `aa * bb`, 1000)
}

func BenchmarkFloatVecAdd_expr(b *testing.B) {
	benchmarkFloat64_expr(b, `aa + bb`, 1000)
}
