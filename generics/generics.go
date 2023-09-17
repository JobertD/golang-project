package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// AddInt is a non-generic function to add two integers.
func AddInt(a int, b int) int {
	return a + b
}

/*
// Non-generic function to add two float64 numbers.
func AddFloat(a float64, b float64) float64 {
    return a + b
}
*/

// DemoNonGenerics demonstrates the usage of a non-generic function.
func DemoNonGenerics() {
	result := AddInt(1, 2)
	fmt.Printf("Non-generic result: %+v\n", result)
}

// Add is a generic function to add two values of any type.
func Add[T int | float64](a T, b T) T {
	return a + b
}

// Num is an interface representing numeric types.
type Num interface {
	int | int8 | int16 | float32 | float64
}

// AddTypeInterface is a generic function that uses an interface constraint.
func AddTypeInterface[T Num](a, b T) T {
	return a + b
}

// AddCO is a generic function that uses the constraints.Ordered constraint.
func AddCO[T constraints.Ordered](a T, b T) T {
	return a + b
}

// DemoGeneric demonstrates the usage of a generic function.
func DemoGeneric() {
	result := Add(1, 2) // Adding integers and float64
	fmt.Printf("Generic result: %d\n", result)
}

// UserID is a custom type representing a user ID.
type UserID int

// AddTil is a generic function using the tilde (~) constraint.
func AddTil[T ~int | float64](a T, b T) T {
	return a + b
}

// DemoGenericTilde demonstrates the usage of a generic function with the tilde constraint.
func DemoGenericTilde() {
	a := UserID(1)
	b := UserID(2)
	result := AddTil(a, b)
	fmt.Printf("Generic using tilde result: %+v\n", result)
}
