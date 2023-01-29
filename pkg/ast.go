package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

type nodeType int

const (
	number nodeType = iota
	operator
	lparen
	rparen
)

type AST struct {
	token Token
	left  *AST
	right *AST
}

type Token struct {
	typ nodeType
	val string
	pos int
}

func parseExpr(expression string) (*AST, int) {
	var tokens = tokenize(expression)
	var outputStack []Token
	var operatorStack []Token
	for _, token := range tokens {
		switch token.typ {
		case number:
			outputStack = append(outputStack, token)
		case operator:
			for len(operatorStack) > 0 && precedence(operatorStack[len(operatorStack)-1].val) >= precedence(token.val) {
				outputStack = append(outputStack, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
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
			if !found {
				return nil, token.pos
			}
		}
	}
	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1].typ == lparen || operatorStack[len(operatorStack)-1].typ == rparen {
			return nil, tokens[len(tokens)-1].pos
		}
		outputStack = append(outputStack, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}
	var astStack []*AST
	for _, token := range outputStack {
		if token.typ == number {
			//	n, _ := strconv.Atoi(token.val)
			astStack = append(astStack, &AST{token: token, left: nil, right: nil})
		} else {
			right := astStack[len(astStack)-1]
			astStack = astStack[:len(astStack)-1]
			left := astStack[len(astStack)-1]
			astStack = astStack[:len(astStack)-1]
			astStack = append(astStack, &AST{token: token, left: left, right: right})
		}
	}
	return astStack[0], 0
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
		} else if char == '(' {
			if buf.Len() > 0 {
				if isNumber(buf.String()) {
					tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
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
			tokens = append(tokens, Token{typ: lparen, val: string(char), pos: i})
		} else if char == ')' {
			if buf.Len() > 0 {
				if isNumber(buf.String()) {
					tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
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
			tokens = append(tokens, Token{typ: rparen, val: string(char), pos: i})
		} else {
			return nil
		}
	}
	if buf.Len() > 0 {
		if isNumber(buf.String()) {
			tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
		} else {
			_, err := strconv.ParseFloat(buf.String(), 64)
			if err == nil {
				tokens = append(tokens, Token{typ: number, val: buf.String(), pos: pos})
			} else {
				return nil
			}
		}
	}
	return tokens
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(token string) int {
	switch token {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func prettyPrint(node *AST, indent string) {
	if node == nil {
		return
	}
	if node.token.typ == number {
		fmt.Printf("%s%s\n", indent, node.token.val)
		return
	}

	fmt.Printf("%s%s\n", indent, node.token.val)
	prettyPrint(node.left, indent+"  ")
	prettyPrint(node.right, indent+"  ")
}