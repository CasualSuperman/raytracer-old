package trace

import (
	"raytracer/color"
	"raytracer/view"
)

const (
	// The number of sub-pixels in an image. Antialiasing jitter is disabled id
	// ANTIALIAS = 1.
	ANTIALIAS = 4
)

// Make a pixel using the number of subsamples indicated in ANTIALIAS.
func makeAntiAliasedPixel(m *view.Model, x, y int, image color.Image) {
	total := color.Color{0, 0, 0}

	for i := 0; i < ANTIALIAS; i++ {
		p := color.Color{0, 0, 0}
		base := mapPixToWorld(m, x, y)
		dist := 0.0
		rayTrace(m, base, &p, dist, nil)

		total.R += p.R
		total.G += p.G
		total.B += p.B
	}

	total.Scale(1 / float64(ANTIALIAS))
	total.Cap(1)

	pixelInImage := image.GetPixel(x, y)

	*pixelInImage.R = uint8(total.R * 255)
	*pixelInImage.G = uint8(total.G * 255)
	*pixelInImage.B = uint8(total.B * 255)
}
