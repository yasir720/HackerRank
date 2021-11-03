package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stack []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(6)
	}
	return slice
}

func rectangleSizeCalculator(height, startingPosition, endPosition int) int {
	var rectangleArea int = height * (endPosition - startingPosition)
	return rectangleArea
}


func largestRectangle(h []int) int {
    // set up our variables
	var positions Stack
	var heights Stack
	var largestRectangle int = 0
	// recurisive closure to aid us when the next height is smaller
	var popThat func(n int) int

	// we range through the input building sizes
	for nextPosition, nextHeight := range h {
		currentPosition := positions[len(positions)-1]
		currentheight := positions[len(heights)-1]

		// the case that the next building is larger
		if nextHeight > currentheight {
			positions.Push(nextPosition)
			heights.Push(nextHeight)
		} else if nextHeight < currentheight { // the case that the next building is smaller
			challengerRectangle := rectangleSizeCalculator(currentheight, currentPosition, nextPosition) // calculate area
			if challengerRectangle > largestRectangle { // update largestArea if needed
				largestRectangle = challengerRectangle
			}
			heights.Pop() // pop the heights because the current rectangle can't keep going
			if nextHeight >= cu//check if the new current height is larger or equal to the next height
		}
	}

	return -1

}

func main() {
	buildings := generateSlice(5)

	one := buildings[2]
	fmt.println(one)
	buildings[2] = 10
	fmt.println(one)

	//largestRectangle(buildings)
}