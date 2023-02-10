package pkg

import (
	"math"
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
		ast, _ := parseExpr(test.expr)
		prettyPrint(ast, " ")
		result := evaluate(ast, vars)
		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("For expression %s, expected %f but got %f", test.expr, test.expected, result)
		}
	}
}

func TestEvaluateRecover(t *testing.T) {
	_, err := Evaluate("aa * 3", nil)
	expected := &evaluateError{message: "Cannot evaluate expression. Key 'aa' not found in environment"}
	if err == nil {
		t.Error("Expected error but got nil")
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error '%s' but got '%s'", expected, err)
	}
}

func TestEvaluateMissingOperand(t *testing.T) {
	_, err := Evaluate("aa * ", nil)
	expected := &evaluateError{message: "Unbalanced expression: not enough operands at position 0"}
	if err == nil {
		t.Error("Expected error but got nil")
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error '%s' but got '%s'", expected, err)
	}
}

func TestEvaluateNoEnv(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ast, _ := parseExpr("aa * 3")
	evaluate(ast, nil)
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
	ast, _ := parseExpr("bb * 3")
	evaluate(ast, vars)
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
	ast, _ := parseExpr("aa * 3")
	evaluate(ast, vars)
}

func TestParseVars(t *testing.T) {
	ast, err := parseExpr("aa[1] * (bb - dd) / c")
	if err != nil {
		t.Error("got error")
	}
	prettyPrint(ast, " ")
}

func TestParseExprErr(t *testing.T) {
	tests := []struct {
		expr     string
		expected error
	}{
		{"(1", &parseError{at: 1, message: "unbalanced parenthesis"}},
		{"1)", &parseError{at: 1, message: "unbalanced parenthesis"}},
		{"1* (2", &parseError{at: 4, message: "unbalanced parenthesis"}},
		{"1 * (2 + 3", &parseError{at: 5, message: "unbalanced parenthesis"}},
		{"1 * (2 + 3))", &parseError{at: 11, message: "unbalanced parenthesis"}},
	}
	for _, test := range tests {
		ast, err := parseExpr(test.expr)
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
