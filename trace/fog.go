package trace

import "raytracer/color"
import "raytracer/debug"
import "raytracer/log"

const (
	// Turns fog on and off.
	FOG_ENABLED = false
)

var (
	fog = struct {
		Color     color.Color
		Thickness float64
	// The color of the fog and its density.
	}{color.Color{0, 0, 0}, 0.15}

	// The maximum distance a ray can travel before it only sees fog.
	maxFogDist = 1 / fog.Thickness * 10
)

// Called during the package's initialization. Prints fog information if it is
// being debugged, and scales the color to be normalized to 1.
func init() {
	if debug.FOG {
		log.Println("Max fog Distance:", maxFogDist)
	}
	fog.Color.Scale(1.0 / 255.0)
}

// Adds fog to a pixel.
func addFog(pixel *color.Color, distance float64) {
	if distance >= maxFogDist {
		pixel.R = fog.Color.R
		pixel.G = fog.Color.G
		pixel.B = fog.Color.B
	} else {
		c := fog.Color
		c.Scale(distance / maxFogDist)
		pixel.Scale(1 - (distance / maxFogDist))

		pixel.R += c.R
		pixel.G += c.G
		pixel.B += c.B
	}
}
