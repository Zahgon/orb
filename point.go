package orb

// A Point is a Lon/Lat 2d point.
type Point [2]float64

var _ Pointer = Point{}

// GeoJSONType returns the GeoJSON type for the object.
func (p Point) GeoJSONType() string {
	_ = "STUB: not implemented"

	// Dimensions returns 0 because a point is a 0d object.
	return ""
}

func (p Point) Dimensions() int {
	_ = "STUB: not implemented"

	// Bound returns a single point bound of the point.
	return 0
}

func (p Point) Bound() Bound {
	_ = "STUB: not implemented"

	// Point returns itself so it implements the Pointer interface.
	return *new(Bound)
}

func (p Point) Point() Point {
	_ = "STUB: not implemented"

	// Y returns the vertical coordinate of the point.
	return *new(Point)
}

func (p Point) Y() float64 {
	_ = "STUB: not implemented"

	// X returns the horizontal coordinate of the point.
	return 0
}

func (p Point) X() float64 {
	_ = "STUB: not implemented"

	// Lat returns the vertical, latitude coordinate of the point.
	return 0
}

func (p Point) Lat() float64 {
	_ = "STUB: not implemented"

	// Lon returns the horizontal, longitude coordinate of the point.
	return 0
}

func (p Point) Lon() float64 {
	_ = "STUB: not implemented"

	// Equal checks if the point represents the same point or vector.
	return 0
}

func (p Point) Equal(point Point) bool { _ = "STUB: not implemented"; return false }
