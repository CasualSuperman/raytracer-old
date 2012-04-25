package shapes

import "io"
import "fmt"

type Material struct {
	Ambient, Diffuse, Specular [3]float64
}

func (m Material) String() string {
	return fmt.Sprintf("\t%v\n\t%v\n\t%v\n", m.Ambient,
		m.Diffuse,
		m.Specular)
}

func (m *Material) Read(r io.Reader) error {
	// Read the ambient.
	count, err := fmt.Fscanf(r, "%d %d %d", &m.Ambient[0], &m.Ambient[1],
		&m.Ambient[2])

	for count == 0 && err == nil {
		count, err = fmt.Fscanf(r, "%d %d %d", &m.Ambient[0], &m.Ambient[1],
			&m.Ambient[2])
	}

	if count != 3 {
		return fmt.Errorf("Tried to read ambient, only got %d values.", count)
	}

	if err != nil {
		return err
	}

	// Read the Diffuse
	count, err = fmt.Fscanf(r, "%d %d %d", &m.Diffuse[0], &m.Diffuse[1],
		&m.Diffuse[2])

	for count == 0 && err == nil {
		count, err = fmt.Fscanf(r, "%d %d %d", &m.Diffuse[0], &m.Diffuse[1],
			&m.Diffuse[2])
	}

	if count != 3 {
		return fmt.Errorf("Tried to read diffuse, only got %d values.", count)
	}

	if err != nil {
		return err
	}

	// Read the Specular
	count, err = fmt.Fscanf(r, "%d %d %d", &m.Specular[0], &m.Specular[1],
		&m.Specular[2])

	for count == 0 && err == nil {
		count, err = fmt.Fscanf(r, "%d %d %d", &m.Specular[0], &m.Specular[1],
			&m.Specular[2])
	}

	if count != 3 {
		return fmt.Errorf("Tried to read specular, only got %d values.", count)
	}

	return err
}
