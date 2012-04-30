package trace

import "raytracer/color"
import "raytracer/debug"
import "raytracer/log"

const (
	FOG_ENABLED = false
)

var (
	fog = struct {
		Color color.Color
		Thickness float64
	}{color.Color{0, 0, 0}, 0.15}

	maxFogDist = 1 / fog.Thickness * 10
)

func init() {
	if debug.FOG {
		log.Println("Max fog Distance:", maxFogDist)
	}
	fog.Color.Scale(1.0/255.0)
}

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
