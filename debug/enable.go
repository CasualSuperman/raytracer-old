package debug

const (
	DIFFUSE  = false
	HITS     = false
	IMAGE    = false
	INPUT    = false
	LIGHTS   = false
	RAYTRACE = false
	PIXEL    = false
	PLANES   = false
	SHAPES   = false
	SPHERES  = false

	ANY = DIFFUSE || HITS || IMAGE || INPUT || LIGHTS || RAYTRACE || PIXEL || PLANES || SHAPES || SPHERES
)
