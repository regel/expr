package ast

import (
	"fmt"
	"strconv"
)

func Evaluate(node *AST, env *Env) interface{} {
	if node == nil {
		errorString := "Cannot evaluate expression: encountered nil token"
		panic(errorString)
	}
	if node.token.typ == number {
		value, _ := strconv.ParseFloat(node.token.val, 64)
		return value
	} else if node.token.typ == name {
		if env == nil {
			errorString := fmt.Sprintf("Cannot evaluate expression. Key '%s' not found in environment", node.token.val)
			panic(errorString)
		}
		value, ok := (*env)[node.token.val]
		if !ok {
			errorString := fmt.Sprintf("Cannot evaluate expression. Key '%s' not found in environment", node.token.val)
			panic(errorString)
		}
		if _, ok := value.(float64); ok {
			return value.(float64)
		}
		if _, ok := value.([]float64); ok {
			return value.([]float64)
		}
		if _, ok := value.(float32); ok {
			return value.(float32)
		}
		if _, ok := value.([]float32); ok {
			return value.([]float32)
		}
		errorString := fmt.Sprintf("Unsupported data type '%T' for token '%v'", value, node.token.val)
		panic(errorString)
	} else if node.token.typ == slice {
		if env == nil {
			errorString := fmt.Sprintf("Cannot evaluate expression. Key '%s' not found in environment", node.token.val)
			panic(errorString)
		}
		value, ok := (*env)[node.token.varName]
		if !ok {
			errorString := fmt.Sprintf("Cannot evaluate expression. Key '%s' not found in environment", node.token.varName)
			panic(errorString)
		}
		if _, ok := value.([]float64); ok {
			return value.([]float64)[node.token.varIdx]
		}
		if _, ok := value.([]float32); ok {
			return value.([]float32)[node.token.varIdx]
		}
		errorString := fmt.Sprintf("Unsupported data type '%T' for token '%v'", value, node.token.varName)
		panic(errorString)
	}

	left := Evaluate(node.left, env)
	right := Evaluate(node.right, env)

	switch node.token.val {
	case "+":
		return add(left, right)
	case "-":
		return subtract(left, right)
	case "*":
		return multiply(left, right)
	case "/":
		return divide(left, right)
	}
	return 0
}
