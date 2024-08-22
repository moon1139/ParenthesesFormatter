package main

import (
	"strings"
)

const operators = "+-*/"

func isOperator(r rune) bool {
	return strings.ContainsRune(operators, r)
}

func isNegativeSign(expRunes []rune, k int) bool {
	n := len(expRunes)
	if (k == 0 && expRunes[k] == '-') && (
		k+1 < n && expRunes[k+1] != '(') {
		return true 												// [-5 ... ], expRunes[k] is a negative sign not an operator
	} else if (k > 0 && expRunes[k] == '-') && (expRunes[k-1] == '(' || expRunes[k-1] == '*' || expRunes[k-1] == '/') {
		return true 												// (-5) or A*-5 or A/-5, means expRunes[k:k+2] are negative number
	}
	return false
}

func updateOperators(expRunes []rune, n int) ([]rune, []rune) {
	leftOp := make([]rune, n)										// memorize last operator of target index 
	rightOp := make([]rune, n)										// memorize next operator of target index

	lastOperator := '\x00'											// A+(B*C)/D, leftOp = ['\x00', '\x00', '+', '+', '+', '*', '*', '*', '/'] 
	for i := 0; i < n; i++ {
		leftOp[i] = lastOperator
		if isOperator(expRunes[i]) && !isNegativeSign(expRunes, i) {
			lastOperator = expRunes[i]
		}
	}

	nextOperator := '\x00'											// A+(B*C)/D, leftOp = ['+', '*', '*', '*', '/', '/', '/', '\x00', '\x00'] 
	for i := n - 1; i >= 0; i-- {
		rightOp[i] = nextOperator
		if isOperator(expRunes[i]) && !isNegativeSign(expRunes, i) {
			nextOperator = expRunes[i]
		}
	}

	return leftOp, rightOp
}

func getOperatorMap(currentOperators []rune) map[rune]bool {
	hasOperator := make(map[rune]bool)								// bool map showing what kind of operator we have in parentheses scope 

	for _, op := range currentOperators {
		hasOperator[op] = true
	}

	return hasOperator
}

func formatParentheses(expression string) string {
	expRunes := []rune(expression)
	n := len(expression)

	reserve := make([]bool, n)										// denote each of the rune if we should reserve or discard
	for i := range reserve {
		reserve[i] = true
	}
													
	leftOp, rightOp := updateOperators(expRunes, n)

	indexStack := []int{}											// store index of parentheses
	operatorStacks := [][]rune{}									// store operator inside parentheses							

	for k := 0; k < n; k++ {
		
		if isOperator(expRunes[k]) && !isNegativeSign(expRunes, k) {
			if len(operatorStacks) > 0 {
				currentStack := &(operatorStacks[len(operatorStacks)-1])
				*currentStack = append(*currentStack, expRunes[k])
			}
		} else if expRunes[k] == '(' {
			indexStack = append(indexStack, k)
			operatorStacks = append(operatorStacks, []rune{})
		} else if expRunes[k] == ')' {
			if len(indexStack)==0 || len(operatorStacks)==0 {
				return "Invalid inputs"
			}

			i := indexStack[len(indexStack)-1] 														// '(' index
			indexStack = indexStack[:len(indexStack)-1]
			j := k 																					// ')' index

			lastOperator := leftOp[i]
			nextOperator := rightOp[j]

			// Exclude operators inside inner parentheses
			currentOperators := operatorStacks[len(operatorStacks)-1]								
			operatorStacks = operatorStacks[:len(operatorStacks)-1] 								// pop stack

			hasOperator := getOperatorMap(currentOperators)

			delFlag := false
			if i > 0 && j < n-1 && expRunes[i-1] == '(' && expRunes[j+1] == ')' { 					// ((scope))
				delFlag = true
			}

			if (lastOperator == '-' || lastOperator == '+') && isNegativeSign(expRunes, i+1) {		// -+(-5)
				// delFlag = false // for human readable
			} else if !hasOperator['+'] && !hasOperator['*'] && !hasOperator['-'] && !hasOperator['/'] { // (noop)
				delFlag = true
			} else if lastOperator == '/' { 														// /(scope)
				// delFlag = false // for human readable
			} else if lastOperator == '-' && (hasOperator['+'] || hasOperator['-']) { 				// -( + or - )
				// delFlag = false // for human readable
			} else if !hasOperator['-'] && !hasOperator['+'] { 										// (* or /)
				delFlag = true
			} else if (i > 0 && j < n-1) && (
				expRunes[i-1] == '(' || lastOperator == '+' || lastOperator == '-') && (
				expRunes[j+1] == ')' || nextOperator == '+' || nextOperator == '-') { 				// ((scope) AND (scope))
				delFlag = true
			} else if (lastOperator == '\x00' || lastOperator == '+' || lastOperator == '-') && (	// noop, +, - (scope)
				nextOperator == '\x00' || nextOperator == '+' || nextOperator == '-') { 			// (scope) noop, +, -
				delFlag = true
			}
			if delFlag {
				reserve[i] = false
				reserve[j] = false
			}
		}
	}

	var res strings.Builder
	for i := 0; i < n; i++ {
		if reserve[i] {
			res.WriteRune(expRunes[i])
		}
	}
	return res.String()
}
