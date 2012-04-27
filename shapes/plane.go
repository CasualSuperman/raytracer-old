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
	RegisterShapeFormat(14, readPlane)
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
		log.Println("Loading Plane normal")
	}

	err = p.Normal.Read(r)

	if err != nil {
		return nil, err
	}

	if debug.PLANES {
		log.Println("Loading Plane center")
	}

	err = p.Center.Read(r)

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

func (p *Plane) Type() shapeId {
	return 14
}

func (p *Plane) Hits(r vector.Ray) (hit bool, length float64, spot *vector.Ray) {
	// Plane normal dot ray direction
	spot = &vector.Ray{p.Center, p.Normal}
	Q := p.Center
	N := p.Normal
	D := r.Direction
	V := r.Position
	P := vector.Origin()
	n_dot_d := vector.Dot(&N, &D)
	T := 0.0

	if debug.PLANES {
		log.Println("Shooting ray starting at", r.Position, "heading", r.Direction, "at plane with center", p.Center, "and normal", p.Normal)
		log.Println("n dot d =", n_dot_d)
	}

	if vector.IsZero(n_dot_d) {
		if debug.PLANES {
			log.Printf("Plane is parallel (n dot d = %f).\n", n_dot_d)
		}
		return false, length, spot
	}

	n_dot_q := vector.Dot(&N, &Q)
	n_dot_v := vector.Dot(&N, &V)

	T = (n_dot_q - n_dot_v) / n_dot_d

	length = T
	/*
		if debug.PLANES {
			log.Printf("N = %s\n", N.String())
			log.Printf("V = %s\n", V.String())
			log.Printf("Q = %s\n", Q.String())
			log.Printf("D = %s\n", D.String())
			log.Printf("nq - nv / nd = %f - %f / %f = %f\n", n_dot_q, n_dot_v, n_dot_d, T)
		}
	*/
	P = D.Position().Copy()
	D2 := P.Direction()
	D2.Scale(T)
	P2 := V.Copy()
	P2.Displace(*D2)

	spot.Position = P2

	if spot.Position.Z > 0 && !vector.IsZero(spot.Position.Z) || length < 0 && !vector.IsZero(length) {
		if debug.PLANES {
			log.Printf("Plane is behind viewer. T = %f, Z position = %f\n", length, spot.Position.Z)
		}
		return false, length, spot
	}

	if debug.PLANES {
		log.Printf("Hit plane %d at point %s, (T = %f)\n", p.Id(), spot.Position.String(), length)
	}

	spot.Direction = p.Normal

	return true, length, spot
}

func (p *Plane) Ambient(d *vector.Position) vector.Vec3 {
	return p.shape.Mat.Ambient
}

func (p *Plane) Diffuse(d *vector.Position) vector.Vec3 {
	return p.shape.Mat.Diffuse
}

func (p *Plane) Specular(d *vector.Position) vector.Vec3 {
	return p.shape.Mat.Specular
}

func (p *Plane) String() string {
	return fmt.Sprintf("Plane:\n\t%v\n\tcenter:\n\t%v\n\tnormal:\n\t%v",
		p.shape.String(), p.Center.String(), p.Normal.String())
}
