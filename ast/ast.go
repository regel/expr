package ast

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type nodeType int

const (
	number nodeType = iota
	operator
	lparen
	rparen
	name
	slice
	function
)

type AST struct {
	token Token
	left  *AST
	right *AST
}

type Token struct {
	typ     nodeType
	val     string
	pos     int
	varName string
	varIdx  int
}

var (
	g_name_pattern *regexp.Regexp
)

func init() {
	g_name_pattern = regexp.MustCompile(`\w+(\[\d+\])?`)
}

func ParseExpr(expression string) (*AST, error) {
	var tokens = tokenize(expression)
	var outputStack []Token
	var operatorStack []Token
	for _, token := range tokens {
		switch token.typ {
		case number:
			outputStack = append(outputStack, token)
		case name:
			if strings.Contains(token.val, "[") {
				i := strings.Index(token.val, "[")
				j := strings.Index(token.val, "]")
				if j == -1 {
					return nil, &ParseError{at: token.pos, message: "Unbalanced expression: missing ']'"}
				}
				if idx, err := strconv.Atoi(token.val[i+1 : j]); err == nil {
					outputStack = append(outputStack, Token{typ: slice, varName: token.val[:i], varIdx: idx})
					continue
				}
				errorString := fmt.Sprintf("Invalid slice index '%s'", token.val[i+1:j])
				return nil, &ParseError{at: token.pos, message: errorString}
			} else {
				outputStack = append(outputStack, token)
			}
		case operator:
			for len(operatorStack) > 0 && precedence(operatorStack[len(operatorStack)-1].val) >= precedence(token.val) {
				outputStack = append(outputStack, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		case function:
			operatorStack = append(operatorStack, token)
		case lparen:
			operatorStack = append(operatorStack, token)
		case rparen:
			found := false
			for len(operatorStack) > 0 {
				if operatorStack[len(operatorStack)-1].typ == lparen {
					found = true
					operatorStack = operatorStack[:len(operatorStack)-1]
					break
				}
				outputStack = append(outputStack, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) > 0 {
				if operatorStack[len(operatorStack)-1].typ == function {
					outputStack = append(outputStack, operatorStack[len(operatorStack)-1])
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
			}
			if !found {
				return nil, &ParseError{at: token.pos, message: "unbalanced parenthesis"}
			}
		}
	}
	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1].typ == lparen || operatorStack[len(operatorStack)-1].typ == rparen {
			return nil, &ParseError{at: tokens[len(tokens)-1].pos, message: "unbalanced parenthesis"}
		}
		outputStack = append(outputStack, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}
	var astStack []*AST
	for _, token := range outputStack {
		if token.typ == number || token.typ == name || token.typ == slice {
			astStack = append(astStack, &AST{token: token, left: nil, right: nil})
		} else if token.typ == function {
			right := astStack[len(astStack)-1]
			astStack = astStack[:len(astStack)-1]
			astStack = append(astStack, &AST{token: token, left: nil, right: right})
		} else {
			right := astStack[len(astStack)-1]
			astStack = astStack[:len(astStack)-1]
			left := astStack[len(astStack)-1]
			astStack = astStack[:len(astStack)-1]
			astStack = append(astStack, &AST{token: token, left: left, right: right})
		}
	}
	return astStack[0], nil
}

func tokenize(expression string) []Token {
	var tokens []Token
	var buf strings.Builder
	var pos int
	for i, char := range expression {
		if char == ' ' {
			continue
		} else if isOperator(string(char)) {
			if buf.Len() > 0 {
				if isNumber(buf.String()) {
					tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
				} else if isFunction(buf.String()) {
					tokens = append(tokens, Token{typ: function, val: buf.String(), pos: pos})
				} else if isName(buf.String()) {
					tokens = append(tokens, Token{typ: name, val: buf.String(), pos: pos})
				} else {
					_, err := strconv.ParseFloat(buf.String(), 64)
					if err == nil {
						tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
					} else {
						return nil
					}
				}
				buf.Reset()
			}
			tokens = append(tokens, Token{typ: operator, val: string(char), pos: i})
		} else if isNumber(string(char)) || string(char) == "." {
			buf.WriteRune(char)
			if pos == 0 {
				pos = i
			}
		} else if isAlpha(char) {
			buf.WriteRune(char)
			if pos == 0 {
				pos = i
			}
		} else if char == '(' {
			if buf.Len() > 0 {
				if isNumber(buf.String()) {
					tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
				} else if isFunction(buf.String()) {
					tokens = append(tokens, Token{typ: function, val: buf.String(), pos: pos})
				} else if isName(buf.String()) {
					tokens = append(tokens, Token{typ: name, val: buf.String(), pos: pos})
				} else {
					_, err := strconv.ParseFloat(buf.String(), 64)
					if err == nil {
						tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
					} else {
						errorString := fmt.Sprintf("found unexpected char '%s' at index %d", string(char), i)
						panic(errorString)
					}
				}
				buf.Reset()
			}
			tokens = append(tokens, Token{typ: lparen, val: string(char), pos: i})
		} else if char == ')' {
			if buf.Len() > 0 {
				if isNumber(buf.String()) {
					tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
				} else if isName(buf.String()) {
					tokens = append(tokens, Token{typ: name, val: buf.String(), pos: pos})
				} else {
					_, err := strconv.ParseFloat(buf.String(), 64)
					if err == nil {
						tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
					} else {
						errorString := fmt.Sprintf("found unexpected char '%s' at index %d", string(char), i)
						panic(errorString)
					}
				}
				buf.Reset()
			}
			tokens = append(tokens, Token{typ: rparen, val: string(char), pos: i})
		} else {
			errorString := fmt.Sprintf("found unexpected char '%s' at index %d", string(char), i)
			panic(errorString)
		}
	}
	if buf.Len() > 0 {
		if isNumber(buf.String()) {
			tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
		} else if isFunction(buf.String()) {
			tokens = append(tokens, Token{typ: function, val: buf.String(), pos: pos})
		} else if isName(buf.String()) {
			tokens = append(tokens, Token{typ: name, val: buf.String(), pos: pos})
		} else {
			_, err := strconv.ParseFloat(buf.String(), 64)
			if err == nil {
				tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
			} else {
				errorString := "found unexpected trailing chars"
				panic(errorString)
			}
		}
	}
	return tokens
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '[' || c == ']') || (c >= '0' && c <= '9')
}

func isName(token string) bool {
	return g_name_pattern.MatchString(token)
}

func isFunction(token string) bool {
	functions := map[string]bool{
		"add":         true,
		"sub":         true,
		"mul":         true,
		"div":         true,
		"min":         true,
		"max":         true,
		"sum":         true,
		"abs":         true,
		"acos":        true,
		"acosh":       true,
		"asin":        true,
		"asinh":       true,
		"atan":        true,
		"atanh":       true,
		"cbrt":        true,
		"ceil":        true,
		"cos":         true,
		"cosh":        true,
		"erf":         true,
		"erfc":        true,
		"erfcinv":     true,
		"erfinv":      true,
		"exp":         true,
		"exp2":        true,
		"expm1":       true,
		"floor":       true,
		"gamma":       true,
		"j0":          true,
		"j1":          true,
		"log":         true,
		"log10":       true,
		"log1p":       true,
		"log2":        true,
		"logb":        true,
		"round":       true,
		"roundToEven": true,
		"sin":         true,
		"sinh":        true,
		"sqrt":        true,
		"tan":         true,
		"tanh":        true,
		"trunc":       true,
		"y0":          true,
		"y1":          true,
		"nanmin":      true,
		"nanmax":      true,
		"nanmean":     true,
		"nanstd":      true,
		"nansum":      true,
		"nanprod":     true,
	}
	return functions[token]
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == ","
}

func precedence(token string) int {
	switch token {
	case "+", "-":
		return 1
	case "*", "/", ",":
		return 2
	}
	return 0
}

func PrettyPrint(w io.Writer, node *AST, indent string) {
	if node == nil {
		return
	}
	if node.token.typ == number {
		fmt.Fprintf(w, "%s%s\n", indent, node.token.val)
		return
	} else if node.token.typ == slice {
		fmt.Fprintf(w, "%s%s[%d]\n", indent, node.token.varName, node.token.varIdx)
		return
	}

	fmt.Fprintf(w, "%s%s\n", indent, node.token.val)
	PrettyPrint(w, node.left, indent+"  ")
	PrettyPrint(w, node.right, indent+"  ")
}

type ParseError struct {
	at      int
	message string
}

func (e *ParseError) Error() string {
	return e.message + " at position " + strconv.Itoa(e.at)
}
