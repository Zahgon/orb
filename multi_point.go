package orb

// A MultiPoint represents a set of points in the 2D Euclidean or Cartesian plane.
type MultiPoint []Point

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPoint) GeoJSONType() string { _ = "STUB: not implemented"; return "" }

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPoint) Dimensions() int {
	_ = "STUB: not implemented"

	// Clone returns a new copy of the points.
	return 0
}

func (mp MultiPoint) Clone() MultiPoint { _ = "STUB: not implemented"; return *new(MultiPoint) }

// Bound returns a bound around the points. Uses rectangular coordinates.
func (mp MultiPoint) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Equal compares two MultiPoint objects. Returns true if lengths are the same
// and all points are Equal, and in the same order.
func (mp MultiPoint) Equal(multiPoint MultiPoint) bool { _ = "STUB: not implemented"; return false }
