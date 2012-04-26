package shapes

import "bufio"
import "fmt"
import "raytracer/vector"

type Material struct {
	Ambient, Diffuse, Specular vector.Vec3
}

func (m Material) String() string {
	return fmt.Sprintf("\t%v\n\t%v\n\t%v",
		m.Ambient, m.Diffuse, m.Specular)
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
