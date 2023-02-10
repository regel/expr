package pkg

import (
	"errors"
	"fmt"
	"strconv"
)

type Env map[string]any

type evaluateError struct {
	message string
}

func (e *evaluateError) Error() string {
	return e.message
}

func Evaluate(e string, env *Env) (v float64, err error) {
	err = nil
	v = 0
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = &evaluateError{
					message: x,
				}
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	node, err := parseExpr(e)
	if err != nil {
		return v, err
	}
	v = evaluate(node, env)
	return v, err
}

func evaluate(node *AST, env *Env) float64 {
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
		errorString := fmt.Sprintf("Unsupported data type '%v' for token '%v'", value, node.token.val)
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
		errorString := fmt.Sprintf("Unsupported data type '%v' for token '%v'", value, node.token.varName)
		panic(errorString)
	}

	left := evaluate(node.left, env)
	right := evaluate(node.right, env)

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
