package shapes

import "bufio"
import "fmt"
import "raytracer/debug"
import "raytracer/log"
import "raytracer/vector"

type Plane struct {
	shape
	Center vector.Position
	Normal vector.Direction
}

func init() {
	RegisterFormat(14, readPlane)
}

func readPlane(r *bufio.Reader) (Shape, error) {
	if debug.PLANES {
		log.Println("Reading in a plane.")
	}
	p := new(Plane)
	err := p.shape.Read(r)
	if err != nil {
		return nil, err
	}

	if debug.PLANES {
		log.Println("Loading Plane center")
	}

	err = p.Center.Read(r)

	if err != nil {
		return nil, err
	}

	if debug.PLANES {
		log.Println("Loading Plane normal")
	}

	err = p.Normal.Read(r)

	if debug.PLANES {
		if err == nil {
			log.Println(p.String())
		} else {
			log.Println("Could not read plane.")
		}
	}

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Plane) Hits(r vector.Ray) (hit bool, length float64, spot vector.Ray) {
	// Plane normal dot ray direction
	nd := vector.Dot(&p.Normal, &r.Direction)

	if vector.IsZero(nd) {
		if debug.PLANES {
			log.Println("Plane is parallel.")
		}
		return false, length, spot
	}

	length = (vector.Dot(&p.Normal, &p.Center) - vector.Dot(&p.Normal, &r.Position)) / nd

	spot.Position = *(r.Direction.Position())
	spot.Position.Direction().Scale(length)
	spot.Position.Offset(r.Position)

	if spot.Position.Z > 0 && !vector.IsZero(nd) {
		if debug.PLANES {
			log.Printf("Plane is behind viewer, (T = %f)\n", length)
		}
		return false, length, spot
	}

	if debug.PLANES {
		log.Printf("Hit plane %d at point %s, (T = %f)\n", p.shape.Id, spot.Position, length)
	}

	spot.Direction = r.Direction

	return true, length, spot
}

func (p *Plane) String() string {
	return fmt.Sprintf("Plane:\n\t%v\n\tcenter:\n\t%v\n\tnormal:\n\t%v",
						p.shape.String(), p.Center.String(), p.Normal.String())
}
