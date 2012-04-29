package debug

const (
	COLOR    = !true
	DIFFUSE  = false
	FPLANES  = false
	HITS     = false
	IMAGE    = false
	INPUT    = false
	LIGHTS   = false
	RAYTRACE = !true
	PIXEL    = false
	PLANES   = false
	SHAPES   = false
	SPECULAR = !true
	SPOTLIGHTS = false
	SPHERES  = false
	TPLANES  = false

	ANY = COLOR || DIFFUSE || FPLANES || HITS || IMAGE || INPUT || LIGHTS || RAYTRACE || PIXEL || PLANES || SPECULAR || SHAPES || SPHERES || TPLANES
)
