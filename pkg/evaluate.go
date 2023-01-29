package pkg

import (
	"strconv"
)

func evaluate(node *AST) int {
	if node == nil {
		return 0
	}
	if node.token.typ == number {
		value, _ := strconv.Atoi(node.token.val)
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
