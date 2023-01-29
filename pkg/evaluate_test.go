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
