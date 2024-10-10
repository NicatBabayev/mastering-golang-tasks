package structs

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) String() string {
	return fmt.Sprintf(" Width: %.1f, Height: %.1f\n Area: %.1f\n Perimeter: %.1f\n", r.Width, r.Height, r.Height*r.Width, 2*(r.Height+r.Width))
}
