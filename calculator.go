//
// Illiasse Mounjim - 102277520   - Oct 19
// Calculator Program.
// able to do addition, subtraction, multiplication and division of whole and float numbers
// uses recursion for parenthesis and calls evaluate and passes what is inside as the expression
//
package main

import (
	"bufio"
	"bytes"
	"math"
)
import "fmt"
import "os"

import "Program2/stack"

// Global operator and operand stacks
var operandStack stack.Stack
var operatorStack stack.Stack



// Returns x ^ y. This is a brute force integer power routine using successive
// multiplication. (There are more efficient ways to do this.)

// Returns true if the character is a digit.
func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// Returns the precedence of the operator.
func precedence(op byte) (prec int) {
	switch op {
	case '+', '-': prec = 0
	case '*', '/': prec = 1
	case '^': prec = 2
	//case '(', ')' : prec = 3
	default: panic("unknown operator")
	}
	return
}

// Returns true if op is right associative. Only exponentiation is right
// associative.
func isRightAssociative(op byte) bool {
	return op == '^'
}

// Add two numbers and return the sum.
func add(x interface{}, y interface{}) interface{} {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int: return x.(int) + y.(int)
		case float64: return float64(x.(int)) + y.(float64)
		}
	case float64:
		switch y.(type) {
		case int: return x.(float64) + float64(y.(int))
		case float64: return x.(float64) + y.(float64)
		}
	}
	panic("unexpected type")
}

// subtract two numbers and return the difference.
func subtract(x interface{}, y interface{}) interface{} {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int: return x.(int) - y.(int)
		case float64: return float64(x.(int)) - y.(float64)
		}
	case float64:
		switch y.(type) {
		case int: return x.(float64) - float64(y.(int))
		case float64: return x.(float64) - y.(float64)
		}
	}
	panic("unexpected type")
}

// multiply two numbers and return the product.
func multiply(x interface{}, y interface{}) interface{} {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int: return x.(int) * y.(int)
		case float64: return float64(x.(int)) * y.(float64)
		}
	case float64:
		switch y.(type) {
		case int: return x.(float64) * float64(y.(int))
		case float64: return x.(float64) * y.(float64)
		}
	}
	panic("unexpected type")
}

// divide two numbers and return the quotient.
func divide(x interface{}, y interface{}) interface{} {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int: return x.(int) / y.(int)
		case float64: return float64(x.(int)) / y.(float64)
		}
	case float64:
		switch y.(type) {
		case int: return x.(float64) / float64(y.(int))
		case float64: return x.(float64) / y.(float64)
		}
	}
	panic("unexpected type")
}

// Returns x ^ y for x and y both integers. This is a brute force integer power
// routine using successive multiplication. (There are more efficient ways to do
// this.)
func intToIntPower(x int, y int) (pow int) {
	pow = 1
	for i := 0 ; i < y ; i++ {
		pow *= x
	}
	return
}
// Returns x ^ y for x float and y integer. This is a brute force integer power
// routine using successive multiplication. (There are more efficient ways to do
// this.)
func floatToIntPower(x float64, y int) (pow float64) {
	pow = 1.0
	for i := 0 ; i < y ; i++ {
		pow *= x
	}
	return
}

// Returns x ^ y.
func power(x interface{}, y interface{}) interface{} {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int: return intToIntPower(x.(int), y.(int))
		case float64: return math.Pow(float64(x.(int)), y.(float64))
		}
	case float64:
		switch y.(type) {
		case int: return floatToIntPower(x.(float64), y.(int))
		case float64: return math.Pow(x.(float64), y.(float64))
		}
	}
	panic("unexpected type")
}


// Apply the top operator on the operator stack to the top two operands on the
// operand stand and push the result onto the operand stack.
func apply(operandStack *stack.Stack, operatorStack *stack.Stack) {
	// Pop the operator off the operator stack
	op, err := operatorStack.Pop()
	if err != nil {
		panic("operator stack underflow")
	}
	// Pop the right operand off the operand stack
	right, err := operandStack.Pop()
	if err != nil {
		panic("operand stack underflow")
	}
	// Pop the left operand off the operand stack
	left, err := operandStack.Pop()
	if err != nil {
		panic("operand stack underflow")
	}
	// Apply the operator to the left and right operands and push the result
	// onto the operand stack
	//switch op.(byte) {

	// Apply the operator to the left and right operands and push the result
	// onto the operand stack
	switch op.(byte) {
	case '+': operandStack.Push(add(left, right))
	case '-': operandStack.Push(subtract(left, right))
	case '*': operandStack.Push(multiply(left, right))
	case '/': operandStack.Push(divide(left, right))
	case '^': operandStack.Push(power(left, right))

	default: panic("unknown operator")
	}
}

// I couldn't get this part running with the applynow.
//I kept getting the following message: interface conversion: interface {} is nil, not uint8
//// Returns true if op1 should be applied before op2.
//func applyNow(op1 byte, op2 byte) bool {
//	if op1 == op2 {
//		return !isRightAssociative(op1)
//	}
//	if precedence(op1) > precedence(op2) {
//		return true
//	}
//	return false
//}

// IsBalanced return true if the expression with parenthesis is balanced return true if
func IsBalanced(expr string) bool {
	isBalanced := true
	s := make([]rune, 0, len(expr))
	for _, c := range expr {
		if c == '(' {
			s = append(s, c)
		} else if c == ')' {
			if len(s) == 0 {
				isBalanced = false
				break
			}
			s = s[:len(s)-1]
		}
	}
	if len(s) != 0 {
		isBalanced = false
	}
	return isBalanced
}


// Evaluate an expression and print the result.
func evaluate(expr string, operandStack *stack.Stack, operatorStack *stack.Stack) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("illegal expression:", r)
		}
	}()


	//checking if an expression is balanced
	if !IsBalanced(expr) {
		fmt.Println("illegal expression: Parenthesis mismatch", expr)
		return
	}

	// Process the expression character by character left to right
	operandExpected := true
	//parenthExpected := true
	i := 0
	for i < len(expr) {
		switch expr[i] {
		// Digit: Extract the operand and push it on the operand stack
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			if !operandExpected {
				panic("operator expected but operand found")
			}
			// Extract whole part
			whole := 0
			for i < len(expr) && isDigit(expr[i]) {
				whole = 10*whole + int(expr[i]-'0')
				i++
			}
			// Extract fractional part
			if i < len(expr) && expr[i] == '.' {
				i++ // Skip over decimal point
				frac := 0.0
				denom := 1.0
				for i < len(expr) && isDigit(expr[i]) {
					denom *= 10.0
					frac += float64(expr[i]-'0')/denom
					i++
				}
				operandStack.Push(float64(whole) + frac)
			} else {
				operandStack.Push(whole)
			}
			operandExpected = false
			//parenthExpected = false

		case '+', '-', '*', '/', '^':
			if operandExpected {
				panic("operand expected but operator found")
			}
			for !operatorStack.IsEmpty() {

				// I couldn't get this part running with the applynow.
				//I kept getting the following message: interface conversion: interface {} is nil, not uint8
				//this was my code with the apply now
				//op1, _ := operatorStack.Pop()
				//if !operatorStack.IsEmpty() {
				//	operatorStack.Push(op1)
				//	break
				//}
				//op2, _ := operatorStack.Top()
				////operatorStack.Push(op1)	//putting operator 1 back on the stack
				//if applyNow(op1.(byte), op2.(byte)) {
				//	apply(operandStack, operatorStack)
				//} else {
				//	break
				//}
				op, _ := operatorStack.Top()
				if precedence(op.(byte)) > precedence(expr[i]) ||
					(precedence(op.(byte)) == precedence(expr[i]) && !isRightAssociative(op.(byte))) {
					apply(operandStack, operatorStack)
				} else {
					break
				}
			}
			operatorStack.Push(expr[i])
			i ++
			operandExpected = true

		case '(':
			var str bytes.Buffer
			i++ //to skip the opening bracket.
			//if we find another opening bracket, we just skip to the next
			//if we find a
			for i< len(expr) && expr[i]!= ')'{
				if byte(expr[i]) != '(' {
					str.WriteString(string(expr[i]))
				}
				i++
			}
			//operand expected after an opening bracket
			operandExpected = true

			if str.Len()!= 0 {
				var TempOperandStack = stack.New() //new stacks for the recursive call.
				var TempOperatorStack = stack.New()
				evaluate(str.String(), &TempOperandStack, &TempOperatorStack)
				var a, _ = TempOperandStack.Pop()
				operandStack.Push(a)
			}
			i++ //to skip the closing bracket.
			operandExpected = false

		case ')':
			operandExpected = false
			i++	//skip the closing bracket.

		case ' ':
			i++
		default:
			panic(fmt.Sprintf("%q is an illegal character", expr[i]))
		}
	}
	// Apply any remaining operators
	for !operatorStack.IsEmpty() {
		apply(operandStack , operatorStack )
	}
	// The result is the one operator remaining on the stack.
	result, _ := operandStack.Pop()
	if !operandStack.IsEmpty() {
		panic("too many operands")
	}
	//fmt.Printf("%v\n", result)
	operandStack.Push(result)

}

// Main routine to read expressions from standard input, calculate their values,
// and print the result. (Use an end of file, control-Z, to exit.)
func main() {
	// Make a scanner to read lines from standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an expression : ")
	//text, _ := scanner.
	// Process each of the lines from standard input
	for scanner.Scan() {
		//fmt.Println("Please enter an expression")

		// Initialize the operator and operand stacks
		operandStack = stack.New()
		operatorStack = stack.New()

		// Get the current line of text.
		line := scanner.Text()
		//fmt.Println(line)

		// Evaluate the expression and print the result
		evaluate(line, &operandStack , &operatorStack)
		result, _ := operandStack.Pop()
		if result != nil {
			fmt.Printf("%v\n", result)
		}
	}
	if scanner.Err() != nil {
		// Handle error.
		fmt.Println("Please enter a valid expression")
	}
}
