package pkg

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expr     string
		expected int
	}{
		{"2 + 3", 5},
		{"2 - 3", -1},
		{"2 * 3", 6},
		{"8 / 4", 2},
		{"2 + 3 * 4", 14},
		{"2 * 3 + 4 * 5", 26},
		{"2 * (3 + 4) * 5", 70},
	}

	for _, test := range tests {
		ast, _ := parseExpr(test.expr)
		prettyPrint(ast, " ")
		result := evaluate(ast)
		if result != test.expected {
			t.Errorf("For expression %s, expected %d but got %d", test.expr, test.expected, result)
		}
	}
}
