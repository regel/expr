package ast

import (
	"fmt"
)

type Args [][]float64

func concat(left, right interface{}) Args {
	if _, ok := left.([]float64); !ok {
		errorString := fmt.Sprintf("Unsupported data type '%T'", left)
		panic(errorString)
	}
	if _, ok := right.([]float64); !ok {
		errorString := fmt.Sprintf("Unsupported data type '%T'", right)
		panic(errorString)
	}

	out := make(Args, 0)
	out = append(out, left.([]float64))
	out = append(out, right.([]float64))
	return out
}
