package main

import (
	"fmt"
	"math"
)

// An interface defines a set of method signatures that a type must implement
// Define the Shape interface with a single method: Area()
type Shape interface {
	Area() float64
}

// Create a struct for Circle
type Circle struct {
	Radius float64
}

// Implement the Area() method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Create a struct for Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area() method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 3.0}

	// Calculate and display the area of the Circle
	fmt.Printf("Circle Area: %.2f\n", circle.Area())

	// Calculate and display the area of the Rectangle
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
}
