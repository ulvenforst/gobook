// Madelbrot emits a PNG image of the Mandelbrot fractal.
package chap3

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func MandelbrotPlot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := range height {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := range width {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrast*n}
		}
	}
	return color.Black
}
