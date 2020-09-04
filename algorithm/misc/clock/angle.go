package clock

import (
	"fmt"
	"math"
)

func Angle(h, m float64) float64 {
	// validate the input
	if h < 0 || m < 0 || h > 12 || m > 60 {
		fmt.Println("Wrong input")
	}

	if h == 12 {
		h = 0
	}

	if m == 60 {
		m = 0
		h += 1
		if h > 12 {
			h = h - 12
		}

	}

	// Calculate the angles moved by hour and
	// minute hands with reference to 12:00
	hAngle := (int)(0.5 * (h*60 + m))
	mAngle := (int)(6 * m)

	// Find the difference between two angles
	angle := math.Abs(float64(hAngle - mAngle))

	// smaller angle of two possible angles
	angle = math.Min(360-angle, angle)

	return angle
}
