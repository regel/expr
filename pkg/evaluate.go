package pkg

import (
	"strconv"
)

func evaluate(node *AST) float64 {
	if node == nil {
		return 0
	}
	if node.token.typ == number {
		value, _ := strconv.ParseFloat(node.token.val, 64)
		return value
	}

	left := evaluate(node.left)
	right := evaluate(node.right)

	switch node.token.val {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return 0
}
