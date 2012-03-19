package shapes

import "raytracer/ray"
import "raytracer/vector"

type Intersecter interface {
	func Intersect(ray.Ray r) (vector.Vec3, bool)
}
