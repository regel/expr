# Expr 

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/license/mit-0/)
[![Go Report Card](https://goreportcard.com/badge/github.com/regel/expr)](https://goreportcard.com/report/github.com/regel/expr)
[![Build](https://github.com/regel/expr/actions/workflows/build.yaml/badge.svg)](https://github.com/regel/expr/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/regel/expr/branch/main/graph/badge.svg)](https://codecov.io/gh/regel/expr)

<img
 src="./static/expr.png"
 raw=true
 alt="Expression language for Go"
 width="150" align="right"
/>

**Expr** package provides an engine that can compile and evaluate math expressions. 
An expression is a one-liner that returns a value (mostly, but not limited to, float64 and []float64).
It is designed for simplicity, speed and safety.

The purpose of the package is to allow users to use expressions for fast vector math operations.
It is a perfect candidate for the foundation of a _time series database_ and _machine learning feature engineering_.
The idea is to let configure things in a dynamic way without recompile of a program:

```coffeescript
# Get a new vector from the addition of cos() and sin() of other variables
cos(X) + 0.33 * sin(Y)

# Get a new norm vector from vector X
(X - nanmin(X)) / (nanmax(X) - nanmin(X))
```

## Features

* Seamless integration with Go (no need to redefine types)
* Static typing
  ```go
  out, err := expr.Compile(`name + age`)
  // err: invalid operation + (mismatched types string and int)
  // | name + age
  // | .....^
  ```
* User-friendly error messages.
* Reasonable set of basic operators.
* Dozens of Numpy-like builtin math functions: `abs`, `acos`, `acosh`, `asin`, `asinh`, `atan`, `atanh`, `cbrt`, `ceil`, `cos`, `cosh`, `erf`, `erfc`, `erfcinv`, `erfinv`, `exp`, `exp2`, `expm1`, `floor`, `gamma`, `j0`, `j1`, `log`, `log10`, `log1p`, `log2`, `logb`, `round`, `roundtoeven`, `sin`, `sinh`, `sqrt`, `tan`, `tanh`, `trunc`, `y0`, `y1`, `maximum`, `minimum`, `mod`, `pow`, `remainder`, `nanmin`, `nanmax`, `nanmean`, `nanstd`, `nansum`, `nanprod`.
  ```coffeescript
  2 * (nanmean(Scores) - minimum(Elevation, Temp))
  ```

## Install

```
go get github.com/regel/expr
```

## Examples

```go
package main

import (
	"fmt"
	"github.com/regel/expr"
	"github.com/regel/expr/ast"
)

func main() {
	code := `1 + (sum(Features) / 2)`

	program, err := expr.Compile(code)
	if err != nil {
		panic(err)
	}

	env := &ast.Env{
		"Features": []float64{0.5, 1.0, 1.5},
	}
	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", output)
}
```


## Who is using Expr?

[Add your company too](https://github.com/regel/expr/edit/main/README.md)

## License

[MIT](LICENSE)
