package expr

import (
	"errors"
	"github.com/regel/expr/ast"
)

type EvaluateError struct {
	Message string
}

func (e *EvaluateError) Error() string {
	return e.Message
}

// Compile parses and compiles given input expression to bytecode program.
func Compile(input string) (node *ast.AST, err error) {
	node = nil
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	node, err = ast.ParseExpr(input)
	return node, err
}

func Run(node *ast.AST, env *ast.Env) (v interface{}, err error) {
	err = nil
	v = 0
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = &EvaluateError{
					Message: x,
				}
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	v = ast.Evaluate(node, env)
	return v, err
}

// Evaluate parses, compiles and runs given input.
func Evaluate(input string, env *ast.Env) (v interface{}, err error) {
	err = nil
	v = 0
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = &EvaluateError{
					Message: x,
				}
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	node, err := ast.ParseExpr(input)
	if err != nil {
		return v, err
	}
	v = ast.Evaluate(node, env)
	return v, err
}
