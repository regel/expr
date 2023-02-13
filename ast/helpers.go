// Code generated by cmd/helpers/main.go. DO NOT EDIT.

package ast

import (
	"fmt"
)

func add(a, b interface{}) interface{} {
	switch x := a.(type) {
	case []float32:
		return addVec(a, b)
	case float32:
		switch y := b.(type) {
		case []float32:
			return addVec(a, b)
		case float32:
			return x + y
		case []float64:
			return addVec(a, b)
		case float64:
			return float64(x) + y
		}
	case []float64:
		return addVec(a, b)
	case float64:
		switch y := b.(type) {
		case []float32:
			return addVec(a, b)
		case float32:
			return x + float64(y)
		case []float64:
			return addVec(a, b)
		case float64:
			return x + y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "+", b))
}

func subtract(a, b interface{}) interface{} {
	switch x := a.(type) {
	case []float32:
		return subtractVec(a, b)
	case float32:
		switch y := b.(type) {
		case []float32:
			return subtractVec(a, b)
		case float32:
			return x - y
		case []float64:
			return subtractVec(a, b)
		case float64:
			return float64(x) - y
		}
	case []float64:
		return subtractVec(a, b)
	case float64:
		switch y := b.(type) {
		case []float32:
			return subtractVec(a, b)
		case float32:
			return x - float64(y)
		case []float64:
			return subtractVec(a, b)
		case float64:
			return x - y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "-", b))
}

func multiply(a, b interface{}) interface{} {
	switch x := a.(type) {
	case []float32:
		return multiplyVec(a, b)
	case float32:
		switch y := b.(type) {
		case []float32:
			return multiplyVec(a, b)
		case float32:
			return x * y
		case []float64:
			return multiplyVec(a, b)
		case float64:
			return float64(x) * y
		}
	case []float64:
		return multiplyVec(a, b)
	case float64:
		switch y := b.(type) {
		case []float32:
			return multiplyVec(a, b)
		case float32:
			return x * float64(y)
		case []float64:
			return multiplyVec(a, b)
		case float64:
			return x * y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "*", b))
}

func divide(a, b interface{}) interface{} {
	switch x := a.(type) {
	case []float32:
		return divideVec(a, b)
	case float32:
		switch y := b.(type) {
		case []float32:
			return divideVec(a, b)
		case float32:
			return x / y
		case []float64:
			return divideVec(a, b)
		case float64:
			return float64(x) / y
		}
	case []float64:
		return divideVec(a, b)
	case float64:
		switch y := b.(type) {
		case []float32:
			return divideVec(a, b)
		case float32:
			return x / float64(y)
		case []float64:
			return divideVec(a, b)
		case float64:
			return x / y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "/", b))
}

func addVec(a, b interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		return addVec(repeatFloat32(x, lenVec(b)), b)
	case []float32:
		switch y := b.(type) {
		case float32:
			return addVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return addVecFloat32(x, y)
		case float64:
			return addVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return addVecFloat64(castFloat64(x), y)
		}
	case float64:
		return addVec(repeatFloat64(x, lenVec(b)), b)
	case []float64:
		switch y := b.(type) {
		case float32:
			return addVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return addVecFloat64(x, castFloat64(y))
		case float64:
			return addVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return addVecFloat64(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "+", b))
}

func addVecFloat32(a, b []float32) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float32, outSize)
	for j, _ := range a {
		out[j] = a[j] + b[j]
	}
	return out
}

func addVecFloat64(a, b []float64) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float64, outSize)
	for j, _ := range a {
		out[j] = a[j] + b[j]
	}
	return out
}

func subtractVec(a, b interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		return subtractVec(repeatFloat32(x, lenVec(b)), b)
	case []float32:
		switch y := b.(type) {
		case float32:
			return subtractVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return subtractVecFloat32(x, y)
		case float64:
			return subtractVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return subtractVecFloat64(castFloat64(x), y)
		}
	case float64:
		return subtractVec(repeatFloat64(x, lenVec(b)), b)
	case []float64:
		switch y := b.(type) {
		case float32:
			return subtractVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return subtractVecFloat64(x, castFloat64(y))
		case float64:
			return subtractVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return subtractVecFloat64(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "-", b))
}

func subtractVecFloat32(a, b []float32) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float32, outSize)
	for j, _ := range a {
		out[j] = a[j] - b[j]
	}
	return out
}

func subtractVecFloat64(a, b []float64) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float64, outSize)
	for j, _ := range a {
		out[j] = a[j] - b[j]
	}
	return out
}

func multiplyVec(a, b interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		return multiplyVec(repeatFloat32(x, lenVec(b)), b)
	case []float32:
		switch y := b.(type) {
		case float32:
			return multiplyVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return multiplyVecFloat32(x, y)
		case float64:
			return multiplyVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return multiplyVecFloat64(castFloat64(x), y)
		}
	case float64:
		return multiplyVec(repeatFloat64(x, lenVec(b)), b)
	case []float64:
		switch y := b.(type) {
		case float32:
			return multiplyVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return multiplyVecFloat64(x, castFloat64(y))
		case float64:
			return multiplyVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return multiplyVecFloat64(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "*", b))
}

func multiplyVecFloat32(a, b []float32) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float32, outSize)
	for j, _ := range a {
		out[j] = a[j] * b[j]
	}
	return out
}

func multiplyVecFloat64(a, b []float64) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float64, outSize)
	for j, _ := range a {
		out[j] = a[j] * b[j]
	}
	return out
}

func divideVec(a, b interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		return divideVec(repeatFloat32(x, lenVec(b)), b)
	case []float32:
		switch y := b.(type) {
		case float32:
			return divideVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return divideVecFloat32(x, y)
		case float64:
			return divideVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return divideVecFloat64(castFloat64(x), y)
		}
	case float64:
		return divideVec(repeatFloat64(x, lenVec(b)), b)
	case []float64:
		switch y := b.(type) {
		case float32:
			return divideVec(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return divideVecFloat64(x, castFloat64(y))
		case float64:
			return divideVec(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return divideVecFloat64(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "/", b))
}

func divideVecFloat32(a, b []float32) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float32, outSize)
	for j, _ := range a {
		out[j] = a[j] / b[j]
	}
	return out
}

func divideVecFloat64(a, b []float64) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float64, outSize)
	for j, _ := range a {
		out[j] = a[j] / b[j]
	}
	return out
}

func max(a, b interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		switch y := b.(type) {
		case float32:
			return MaxFloat32(x, y)
		case float64:
			return MaxFloat64(float64(x), y)
		default:
			return max(repeatFloat32(x, lenVec(b)), b)
		}
	case []float32:
		switch y := b.(type) {
		case float32:
			return max(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return maxFloat32(x, y)
		case float64:
			return max(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return maxFloat64(castFloat64(x), y)
		}
	case float64:
		switch y := b.(type) {
		case float32:
			return MaxFloat64(x, float64(y))
		case float64:
			return MaxFloat64(x, y)
		default:
			return max(repeatFloat64(x, lenVec(b)), b)
		}
	case []float64:
		switch y := b.(type) {
		case float32:
			return max(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return maxFloat64(x, castFloat64(y))
		case float64:
			return max(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return maxFloat64(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %v %T %T", "Max", a, b))
}

func maxFloat32(a, b []float32) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float32, outSize)
	for j, _ := range a {
		out[j] = MaxFloat32(a[j], b[j])
	}
	return out
}

func maxFloat64(a, b []float64) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float64, outSize)
	for j, _ := range a {
		out[j] = MaxFloat64(a[j], b[j])
	}
	return out
}

func min(a, b interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		switch y := b.(type) {
		case float32:
			return MinFloat32(x, y)
		case float64:
			return MinFloat64(float64(x), y)
		default:
			return min(repeatFloat32(x, lenVec(b)), b)
		}
	case []float32:
		switch y := b.(type) {
		case float32:
			return min(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return minFloat32(x, y)
		case float64:
			return min(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return minFloat64(castFloat64(x), y)
		}
	case float64:
		switch y := b.(type) {
		case float32:
			return MinFloat64(x, float64(y))
		case float64:
			return MinFloat64(x, y)
		default:
			return min(repeatFloat64(x, lenVec(b)), b)
		}
	case []float64:
		switch y := b.(type) {
		case float32:
			return min(a, repeatFloat32(y, lenVec(a)))
		case []float32:
			return minFloat64(x, castFloat64(y))
		case float64:
			return min(a, repeatFloat64(y, lenVec(a)))
		case []float64:
			return minFloat64(x, y)
		}
	}
	panic(fmt.Sprintf("invalid operation: %v %T %T", "Min", a, b))
}

func minFloat32(a, b []float32) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float32, outSize)
	for j, _ := range a {
		out[j] = MinFloat32(a[j], b[j])
	}
	return out
}

func minFloat64(a, b []float64) interface{} {
	outSize := len(a)
	if len(b) < len(a) {
		outSize = len(b)
	}
	out := make([]float64, outSize)
	for j, _ := range a {
		out[j] = MinFloat64(a[j], b[j])
	}
	return out
}

func abs(a interface{}) interface{} {
	switch x := a.(type) {
	case float32:
		return AbsFloat32(x)
	case []float32:
		return absFloat32(x)
	case float64:
		return AbsFloat64(x)
	case []float64:
		return absFloat64(x)
	}
	panic(fmt.Sprintf("invalid operation: %v %T", "Abs", a))
}

func absFloat32(a []float32) interface{} {
	out := make([]float32, len(a))
	for j, _ := range a {
		out[j] = AbsFloat32(a[j])
	}
	return out
}

func absFloat64(a []float64) interface{} {
	out := make([]float64, len(a))
	for j, _ := range a {
		out[j] = AbsFloat64(a[j])
	}
	return out
}

func castFloat64(a interface{}) []float64 {
	switch x := a.(type) {
	case []float32:
		out := make([]float64, 0)
		for _, val := range x {
			out = append(out, float64(val))
		}
		return out
	}
	panic(fmt.Sprintf("invalid operation: %v %T", "cast float64", a))
}

func repeatFloat32(val float32, length int) []float32 {
	out := make([]float32, length)
	for i := range out {
		out[i] = val
	}
	return out
}

func repeatFloat64(val float64, length int) []float64 {
	out := make([]float64, length)
	for i := range out {
		out[i] = val
	}
	return out
}

func AbsFloat32(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func AbsFloat64(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a

}

func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func lenVec(a interface{}) int {
	switch x := a.(type) {
	case []float32:
		return len(x)
	case []float64:
		return len(x)
	}
	panic(fmt.Sprintf("invalid operation: %v %T", "len", a))
}