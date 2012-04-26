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
	RegisterFormat(13, readSphere)
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
	start := r.Position.Copy()
	start.Displace(start.Offset(s.Center))

	radius := s.Radius

	a := vector.Dot(&r.Direction, &r.Direction)
	b := 2 * vector.Dot(&r.Position, &r.Direction)
	c := vector.Dot(&r.Position, &r.Position) - radius*radius

	discriminant := b*b - 4*a*c

	if math.Abs(discriminant) <= math.SmallestNonzeroFloat32 {
		return false, math.Inf(1), spot
	}

	length = (-b - math.Sqrt(discriminant)) / 2 * a

	move := r.Direction.Copy()
	hitPos := r.Position.Copy()
	move.Scale(length)
	hitPos.Displace(move)

	// Set the hit position to where we hit it.
	spot.Position = hitPos

	// Start the direction at the hit position, and subtract the center of the
	// sphere.
	spot.Direction = *hitPos.Direction()
	spot.Direction.Sub(s.Center.Direction())

	return
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere:\n\t%v\n\tcenter:\n\t%v\n\tradius:\n\t%v",
		s.shape.String(), s.Center.String(), s.Radius)
}
