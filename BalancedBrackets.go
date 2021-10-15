package main

import (
	"fmt"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

func oddOrEvan(n int) bool {
	if n % 2 == 0 {
		return true
	} else {
		return false
	}
}

// Helper function to help determine if braket pairs are a match
func isMatch(n string, k byte) bool {
	if string(n) == "(" && string(k) == ")" {
		return true
	} else if string(n) == "[" && string(k) == "]" {
		return true
	} else if string(n) == "{" && string(k) == "}" {
		return true
	} else {
		return false
	}	
}
func isBalanced(s string) string {
    // Write your code here
	var stack Stack

	if !oddOrEvan(len(s)) {
		fmt.Println("OE")
		return "NO"		
	}

	for _, current := range s {
		var topElement string = ""

		if !stack.IsEmpty() {
			topElement = stack[len(stack)-1]
		}		
		fmt.Println("top", topElement, "current", string(current))
		fmt.Println("is match?", isMatch(topElement, byte(current)))
		
		

		// if topElement == current { // if current is the same as the top element of the stack then we pop (remove)
		// 	stack.Pop()
		// } 
		if isMatch(topElement, byte(current)) {
			stack.Pop()
			fmt.Println("pop operation")
		} else {
			stack.Push(string(current))
			fmt.Println("push operation")
		}
		fmt.Println("stack ==>", stack)
		fmt.Println("--------------------------------------")
		// else if topElement != current { // if current isn't the same as the top element of the stack push (add)
		// 	stack.Push(string(current))
		// }

		

		// fmt.Println("loop")
		// fmt.Println(stack)
	}
	
	// if len(stack) == len(s)/2  {
	// 	fmt.Println("len check end")
	// 	fmt.Println(stack)
	// 	return "YES"
	// }
	// fmt.Println("default")
	// fmt.Println(stack)
	if stack.IsEmpty() {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	//var stack Stack // create a stack variable of type Stack

	// stack.Push("this")
	// stack.Push("is")
	// stack.Push("sparta!!")

	// fmt.Println(stack)

	// for len(stack) > 0 {
	// 	x, y := stack.Pop() // This removes items and makes the length shorter to keep the loop going
	// 	if y == true {
	// 		fmt.Println(x)
	// 	}
	// }
	fmt.Println(isBalanced("{}"))	
	// poop := "({{[]}})"
	// fmt.Println(isMatch(poop[2], poop[4]))
}