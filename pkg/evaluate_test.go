package pkg

import (
	"math"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"2 + 3", 5},
		{"2 - 3", -1},
		{"2 * 3", 6},
		{"8 / 4", 2},
		{"2 + 3 * 4", 14},
		{"2 * 3 + 4 * 5", 26},
		{"2 * (3 + 4) * 5", 70},
		{"2.21 * (3.07 + 4) * 5.001", 78.1391247},
	}

	for _, test := range tests {
		ast, _ := parseExpr(test.expr)
		prettyPrint(ast, " ")
		result := evaluate(ast)
		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("For expression %s, expected %f but got %f", test.expr, test.expected, result)
		}
	}
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
