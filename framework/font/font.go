package font

import (
	"github.com/gop9/fonts/bariol/bariolregular"
)

// Returns default font (bariolregular) with specified size and scaling
func DefaultFont(size int, scaling float64) Face {
	face, err := NewFace(bariolregular.TTF, int(float64(size)*scaling))
	if err != nil {
		panic(err)
	}
	return face
}
