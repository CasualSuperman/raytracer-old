package debug

const (
	DIFFUSE  = false
	HITS     = false
	IMAGE    = false
	INPUT    = false
	LIGHTS   = false
	RAYTRACE = false
	PLANES   = false
	SHAPES   = false
	SPHERES  = false

	ANY = DIFFUSE || HITS || IMAGE || INPUT || LIGHTS || RAYTRACE || PLANES || SHAPES || SPHERES
)
