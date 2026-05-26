package orb

// Polygon is a closed area. The first LineString is the outer ring.
// The others are the holes. Each LineString is expected to be closed
// ie. the first point matches the last.
type Polygon []Ring

// GeoJSONType returns the GeoJSON type for the object.
func (p Polygon) GeoJSONType() string {
	_ = "STUB: not implemented"

	// Dimensions returns 2 because a Polygon is a 2d object.
	return ""
}

func (p Polygon) Dimensions() int {
	_ = "STUB: not implemented"

	// Bound returns a bound around the polygon.
	return 0
}

func (p Polygon) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Equal compares two polygons. Returns true if lengths are the same
// and all points are Equal.
func (p Polygon) Equal(polygon Polygon) bool { _ = "STUB: not implemented"; return false }

// Clone returns a new deep copy of the polygon.
// All of the rings are also cloned.
func (p Polygon) Clone() Polygon { _ = "STUB: not implemented"; return *new(Polygon) }
