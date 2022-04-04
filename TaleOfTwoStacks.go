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

// Pop elements from firstStack to secondStack to make a queue
func susNoMore(firstStack, secondStack *Stack) {
	topElement, hasElements := firstStack.Pop()

	if hasElements {
		secondStack.Push(topElement)
	}
}

func SusQueue(queryType, value int) {
	var firstStack Stack
	var secondStack Stack
	var topElement string = ""

	if queryType == 1 { // Enqueue
		firstStack.Push(fmt.Sprint(value))
	} else if queryType == 2 { // Dequeue
		susNoMore(&firstStack, &secondStack)
	} else if queryType == 3 { // Print first element
		if !secondStack.IsEmpty() {
			topElement = secondStack[len(secondStack)-1]
		}
		fmt.Println(topElement)

	}

	fmt.Println("first stack", firstStack)
	fmt.Println("second stack", secondStack)
}

func main() {
	for i := 0; i < 4; i++ {
		SusQueue(1, i)
	}
	SusQueue(3, 0)
	
	
	//fmt.Println(isBalanced("{}"))
}