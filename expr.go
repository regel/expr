// Package expr provides a framework for parsing and evaluating mathematical expressions.
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

// Compile parses and compiles a given input expression.
// The input expression is parsed into an AST (Abstract Syntax Tree) node, which
// is then returned along with any errors that occurred during the parsing process.
//
// Parameters:
//   - input (string): The input expression to be compiled.
//
// Returns:
//   - node (*ast.AST): The root node of the compiled AST.
//   - err (error): Any error that occurred during parsing, or nil if no errors occurred.
//
// If a panic occurs during parsing, it is caught and an error is returned with a message describing the cause of the panic.
// If the panic is not a string or an error, an "unknown panic" error is returned.
//
// Example usage:
//
// node, err := ast.Compile("1 + 2 * 3")
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

// Run executes the given AST in the provided environment.
//
// Parameters:
//   - node (*ast.AST): the AST to be executed.
//   - env (*ast.Env): the environment to be used for execution.
//
// Return values:
//   - v (interface{}): the result of executing the AST.
//   - err (error): an error, if one occurred during execution.
//
// If a panic occurs during execution, it is caught and an error is returned with a message describing the cause of the panic. If the panic is not a string or an error, an "unknown panic" error is returned.
//
// Example:
//
//	package main
//	import (
//	  "fmt"
//	  "github.com/regel/expr/ast"
//	)
//	func main() {
//	  env := ast.NewEnv()
//	  env.Set("x", 2.0)
//	  env.Set("y", 3.0)
//	  input := "x + y"
//	  node, err := ast.ParseExpr(input)
//	  if err != nil {
//	    fmt.Println(err)
//	  } else {
//	    result, err := expr.Run(node, env)
//	    if err != nil {
//	      fmt.Println(err)
//	    } else {
//	      fmt.Println(result)
//	    }
//	  }
//	}
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

// Evaluate parses and evaluates a given input string using the provided environment.
//
// Parameters:
//   - input (string): the input string to be evaluated.
//   - env (*ast.Env): the environment to be used for evaluating the input string.
//
// Return values:
//   - v (interface{}): the result of evaluating the input string.
//   - err (error): an error, if one occurred during evaluation.
//
// If a panic occurs during evaluation, it is caught and an error is returned with a message describing the cause of the panic.
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
