package geometry

import (
	"math"
	"testing"
)

type MockShape struct {
	area            float64
	perimeter       float64
	areaCalled      bool
	perimeterCalled bool
}

func (m *MockShape) Area() float64 {
	m.areaCalled = true
	return m.area
}

func (m *MockShape) Perimeter() float64 {
	m.perimeterCalled = true
	return m.perimeter
}

func TestPrintShapeInfo(t *testing.T) {
	mockRect := &MockShape{area: 20.0, perimeter: 18.0}
	mockCircle := &MockShape{area: 27.9, perimeter: 18.5}

	printShapeInfo("Rectangle", mockRect)
	printShapeInfo("Circle", mockCircle)

	if !mockRect.areaCalled || !mockRect.perimeterCalled {
		t.Errorf("Expected to call Area and Perimeter methods but not called for Rectangle")
	}

	if !mockCircle.areaCalled || !mockCircle.perimeterCalled {
		t.Errorf("Expected to call Area and Perimeter methods but not called for Circle")
	}
}

func TestGeometry(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{Width: 4, Height: 5}
		checkShape(t, "Rectangle", r)
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{Radius: 3}
		checkShape(t, "Circle", c)
	})
}

func checkShape(t *testing.T, shapeName string, s Shape) {
	t.Parallel()
	t.Run("Area", func(t *testing.T) {
		area := s.Area()
		expectedArea := CalculateExpectedArea(shapeName, s)

		if area != expectedArea {
			t.Errorf("Expected %s area to be %.2f, but got %.2f", shapeName, expectedArea, area)
		}

	})
}

func CalculateExpectedArea(shapeName string, s Shape) float64 {
	switch shapeName {
	case "Rectangle":
		return 20
	case "Circle":
		return math.Pi * 3 * 3
	}
	return 0
}
