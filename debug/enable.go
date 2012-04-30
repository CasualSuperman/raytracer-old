package debug

const (
	COLOR      = false
	CYLINDERS  = false
	DIFFUSE    = false
	FOG        = false
	FPLANES    = false
	HITS       = false
	IMAGE      = false
	INPUT      = false
	LIGHTS     = false
	RAYTRACE   = false
	PIXEL      = false
	PLANES     = false
	PPLANES    = false
	SHAPES     = false
	SPECULAR   = false
	SPOTLIGHTS = false
	SPHERES    = false
	TPLANES    = false

	ANY = COLOR || CYLINDERS || DIFFUSE || FPLANES || HITS || IMAGE || INPUT || LIGHTS || RAYTRACE || PIXEL || PLANES || SPECULAR || SHAPES || SPHERES || TPLANES
)
