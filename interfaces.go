package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
	getCircumference() float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

// square methods
func (square Square) getArea() float64 {
	return square.length * square.length
}

func (square Square) getCircumference() float64 {
	return square.length * 4
}

// circle methods
func (circle Circle) getArea() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle Circle) getCircumference() float64 {
	return 2 * math.Pi * circle.radius
}

func printShapeInfo(shape Shape) {
	fmt.Printf("Area of %T is: %0.2f\n", shape, shape.getArea())
	fmt.Printf("Circumference of %T is: %0.2f\n", shape, shape.getCircumference())
}

func interfacesDemo() {
	shapes := []Shape{
		Square{length: 3},
		Circle{radius: 2},
		Square{length: 6},
		Circle{radius: 10},
	}

	for _, value := range shapes {
		printShapeInfo(value)
		fmt.Println("\t-------")
	}
}
