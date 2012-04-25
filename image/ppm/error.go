package ppm

type FormatError string
type UnsupportedError string

func (e FormatError) Error() string {
	return "ppm: invalid format: " + string(e)
}

func (e UnsupportedError) Error() string {
	return "ppm: unsupported feature: " + string(e)
}
