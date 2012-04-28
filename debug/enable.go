package debug

const (
	DIFFUSE  = false
	FPLANES  = false
	HITS     = false
	IMAGE    = false
	INPUT    = false
	LIGHTS   = false
	RAYTRACE = false
	PIXEL    = false
	PLANES   = false
	SHAPES   = false
	SPHERES  = false
	TPLANES  = false

	ANY = DIFFUSE || FPLANES || HITS || IMAGE || INPUT || LIGHTS || RAYTRACE || PIXEL || PLANES || SHAPES || SPHERES || TPLANES
)
