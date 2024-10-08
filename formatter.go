package main

import (
	"strings"
)

const operators = "+-*/"

func isOperator(r rune) bool {
	return strings.ContainsRune(operators, r)
}

func formatParentheses(expression string) string {
	expRunes := []rune(expression)
	n := len(expression)

	reserve := make([]bool, n) 							// denote each of the rune if we should reserve or discard
	for i := range reserve {
		reserve[i] = true
	}

	leftOp := make([]rune, n)  							// memorize last met operator of target index
	rightOp := make([]rune, n) 							// memorize next operator of target index

	lastOperator := '\x00' 	   							// '\x00' denotes default value for left operator
	for i := 0; i < n; i++ {
		leftOp[i] = lastOperator
		if isOperator(expRunes[i]) {
			lastOperator = expRunes[i]
		}
	}

	nextOperator := '\x00' 	   							// '\x00' denotes default value for right operator
	for i := n - 1; i >= 0; i-- {
		rightOp[i] = nextOperator
		if isOperator(expRunes[i]) {
			nextOperator = expRunes[i]
		}
	}

	indexStack := []int{} 								// store index of parentheses
	lastIndexOfOperator := make(map[rune]int) 			// memorize last occurence of the operator
	hasOperator := make(map[rune]bool) 					// denote if we have operator in our current scope (inside parentheses)
	for _, op := range operators {
		lastIndexOfOperator[op] = -1
	}	

	for k := 0; k < n; k++ {
		if isOperator(expRunes[k]) {
			op:= expRunes[k]
			lastIndexOfOperator[op] = k					// continuously update index of last seen op
		}
		
		if expRunes[k] == '(' {
			indexStack = append(indexStack, k) 			// store '(' index to stack
		} else if expRunes[k] == ')' {					// met   ')', we can start varify
			for _, op := range operators {
				hasOperator[op] = false
			}

			i := indexStack[len(indexStack)-1]  		// left parenthesis index
			indexStack = indexStack[:len(indexStack)-1] // pop stack
			j := k                                      // right parenthesis index

			lastOperator := leftOp[i]
			nextOperator := rightOp[j]

			for _, op := range operators {
				if lastIndexOfOperator[op] >= i {
					hasOperator[op] = true
				}
			}

			delFlag := false
			if i > 0 && j+1 < n && expRunes[i-1] == '(' && expRunes[j+1] == ')' { 		// ((scope))
				delFlag = true
			}
			if !hasOperator['+'] && !hasOperator['*'] && !hasOperator['-'] && !hasOperator['/'] { // (no op here)
				delFlag = true
			}

			if lastOperator == '/' { 													// /(scope)
				// delFlag = false // for human readable
			} else if lastOperator == '-' && (hasOperator['+'] || hasOperator['-']) { 	// -( + or - )
				// delFlag = false // for human readable
			} else if !hasOperator['-'] && !hasOperator['+'] { 							// (* or /)
				delFlag = true
			} else if (lastOperator == '\x00' || lastOperator == '+' || lastOperator == '-') && (
					   nextOperator == '\x00' || nextOperator == '+' || nextOperator == '-') { // noop or + or - (scope) AND (scope) noop or + or -
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