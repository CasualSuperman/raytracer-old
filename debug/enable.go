package debug

const (
	DIFFUSE  = !true
	HITS     = false
	IMAGE    = false
	INPUT    = false
	LIGHTS   = false
	RAYTRACE = false
	PIXEL    = !true
	PLANES   = false
	SHAPES   = false
	SPHERES  = false

	ANY = DIFFUSE || HITS || IMAGE || INPUT || LIGHTS || RAYTRACE || PIXEL || PLANES || SHAPES || SPHERES
)
