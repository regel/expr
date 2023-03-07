package ast

import (
	"fmt"
	"strconv"
)

func Evaluate(node *AST, env *Env) interface{} {
	if node == nil {
		return nil
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
	case ",":
		return concat(left, right)
	case "sum":
		return sum(right)
	case "abs":
		return abs(right)
	case "acos":
		return acos(right)
	case "acosh":
		return acosh(right)
	case "asin":
		return asin(right)
	case "asinh":
		return asinh(right)
	case "atan":
		return atan(right)
	case "atanh":
		return atanh(right)
	case "cbrt":
		return cbrt(right)
	case "ceil":
		return ceil(right)
	case "cos":
		return cos(right)
	case "cosh":
		return cosh(right)
	case "erf":
		return erf(right)
	case "erfc":
		return erfc(right)
	case "erfcinv":
		return erfcinv(right)
	case "erfinv":
		return erfinv(right)
	case "exp":
		return exp(right)
	case "exp2":
		return exp2(right)
	case "expm1":
		return expm1(right)
	case "floor":
		return floor(right)
	case "gamma":
		return gamma(right)
	case "j0":
		return j0(right)
	case "j1":
		return j1(right)
	case "log":
		return log(right)
	case "log10":
		return log10(right)
	case "log1p":
		return log1p(right)
	case "log2":
		return log2(right)
	case "logb":
		return logb(right)
	case "round":
		return round(right)
	case "roundtoeven":
		return roundtoeven(right)
	case "sin":
		return sin(right)
	case "sinh":
		return sinh(right)
	case "sqrt":
		return sqrt(right)
	case "tan":
		return tan(right)
	case "tanh":
		return tanh(right)
	case "trunc":
		return trunc(right)
	case "y0":
		return y0(right)
	case "y1":
		return y1(right)
	case "add":
		args := right.(Args)
		return add(args[0], args[1])
	case "sub":
		args := right.(Args)
		return subtract(args[0], args[1])
	case "mul":
		args := right.(Args)
		return multiply(args[0], args[1])
	case "div":
		args := right.(Args)
		return divide(args[0], args[1])
	case "min":
		args := right.(Args)
		return min(args[0], args[1])
	case "max":
		args := right.(Args)
		return max(args[0], args[1])
	case "nanmin":
		return nanmin(right)
	case "nanmax":
		return nanmax(right)
	case "nanmean":
		return nanmean(right)
	case "nanstd":
		return nanstd(right)
	case "nansum":
		return nansum(right)
	case "nanprod":
		return nanprod(right)
	}
	return 0
}
