package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got: %.2f | Want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// Table Driven Tests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// The test speaks to us more clearly, as if it were an assertion of truth,
		// not a sequence of operations :)

		// {Rectangle{12, 6}, 72.0},
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		// {Circle{10}, 314.1592653589793},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		// {Triangle{12, 6}, 36.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v Got: %g | Want: %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("Got: %g | Want: %g", got, want)
	// 	}
	// }

	// t.Run("Rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("Circles", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 72.0)
	// })
}
