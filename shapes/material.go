package shapes

import "bufio"
import "fmt"
import "raytracer/color"

type Material struct {
	Ambient, Diffuse, Specular color.Color
}

func (m Material) String() string {
	return fmt.Sprintln("\t\tAmbient:", m.Ambient, "\n\t\tDiffuse:", m.Diffuse,
		"\n\t\tSpecular:", m.Specular)
}

func (m *Material) Read(r *bufio.Reader) error {
	err := m.Ambient.Read(r)

	if err != nil {
		return fmt.Errorf("Failed to read in ambient: %v", err)
	}

	err = m.Diffuse.Read(r)

	if err != nil {
		return fmt.Errorf("Failed to read in diffuse: %v", err)
	}

	err = m.Specular.Read(r)

	if err != nil {
		return fmt.Errorf("Failed to read in specular: %v", err)
	}
	return nil
}
