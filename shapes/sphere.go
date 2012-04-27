package shapes

import "bufio"
import "fmt"
import "math"
import "raytracer/debug"
import "raytracer/log"
import "raytracer/vector"

type Sphere struct {
	shape
	Center vector.Position
	Radius float64
}

func init() {
	RegisterShapeFormat(13, readSphere)
}

func readSphere(r *bufio.Reader) (Shape, error) {
	if debug.SPHERES {
		log.Println("Reading in a sphere.")
	}
	s := new(Sphere)
	err := s.shape.Read(r)
	if err != nil {
		return nil, err
	}

	if debug.SPHERES {
		log.Println("Loading Sphere center")
	}

	s.Center.Read(r)

	if debug.SPHERES {
		log.Println("Loading Sphere radius")
	}

	count, line := 0, []byte{}

	for count == 0 && err == nil {
		line, _, err = r.ReadLine()
		if err == nil {
			count, _ = fmt.Sscanf(string(line), "%f", &s.Radius)
		}
	}

	if debug.SPHERES {
		if err == nil {
			log.Println(s.String())
		} else {
			log.Println("Could not read sphere.")
		}
	}

	return s, nil
}

func (s *Sphere) Hits(r vector.Ray) (hit bool, length float64, spot vector.Ray) {
/*
	if debug.SPHERES {
		log.Println("Ray:")
		log.Println("\tPosition:", r.Position)
		log.Println("\tDirection", r.Direction)
		log.Println("Sphere:")
		log.Println("\tCenter:", s.Center)
		log.Println("\tRadius:", s.Radius)
	}
*/

	radius := s.Radius

	newPos := r.Position.Copy()

	offset := s.Center.Copy()
	newPos.Displace(offset.Offset(vector.Origin()))

	if debug.SPHERES {
		log.Println(newPos)
	}

	a := vector.Dot(&r.Direction, &r.Direction)
	b := 2 * vector.Dot(&newPos, &r.Direction)
	c := vector.Dot(&newPos, &newPos) - radius*radius

	if debug.SPHERES {
		log.Printf("a: %f, b: %f, c: %f\n", a, b, c)
	}

	discriminant := b*b - 4*a*c

	if discriminant <= 0 {
		return false, math.Inf(1), spot
	}

	length = (-b - math.Sqrt(discriminant)) / (2 * a)

	if debug.SPHERES {
		log.Printf("Sphere discriminant is %f, (T = %f)\n", discriminant, length)
	}

	move := r.Direction.Copy()
	hitPos := newPos.Copy()
	move.Scale(length)
	hitPos.Displace(move)

	// Set the hit position to where we hit it.
	spot.Position = hitPos

	// Start the direction at the hit position, and subtract the center of the
	// sphere.
	spot.Direction = *hitPos.Direction()
	spot.Direction.Sub(s.Center.Direction())

	hit = true

	return
}

func (s *Sphere) Ambient(p *vector.Position) vector.Vec3 {
	return s.shape.Mat.Ambient
}

func (s *Sphere) Diffuse(p *vector.Position) vector.Vec3 {
	return s.shape.Mat.Diffuse
}

func (s *Sphere) Specular(p *vector.Position) vector.Vec3 {
	return s.shape.Mat.Specular
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere:\n\t%v\n\tcenter:\n\t%v\n\tradius:\n\t%v",
		s.shape.String(), s.Center.String(), s.Radius)
}
