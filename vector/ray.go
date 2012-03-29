package vector

/* Inherited from Direction:
 * 		Scale
 *
 * Inherited from Position:
 * 		Displace
 */
func NewRay(p Position, d Direction) Ray {
	return Ray{p, d}
}
