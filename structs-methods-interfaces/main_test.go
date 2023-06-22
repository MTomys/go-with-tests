package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	expect := 40.0
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)

	if expect != result {
		t.Errorf("expected: %.2f, received: %.2f", expect, result)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"Rectangle", Rectangle{Width: 12, Height: 6}, 72.0},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			received := tt.shape.Area()
			if tt.expected != received {
				t.Errorf("Expected: %g, received: %g", tt.expected, received)
			}

		})
	}
}
