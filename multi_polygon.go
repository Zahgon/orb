package orb

// MultiPolygon is a set of polygons.
type MultiPolygon []Polygon

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPolygon) GeoJSONType() string { _ = "STUB: not implemented"; return "" }

// Dimensions returns 2 because a MultiPolygon is a 2d object.
func (mp MultiPolygon) Dimensions() int {
	_ = "STUB: not implemented"

	// Bound returns a bound around the multi-polygon.
	return 0
}

func (mp MultiPolygon) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Equal compares two multi-polygons.
func (mp MultiPolygon) Equal(multiPolygon MultiPolygon) bool {
	_ = "STUB: not implemented"
	return false
}

// Clone returns a new deep copy of the multi-polygon.
func (mp MultiPolygon) Clone() MultiPolygon { _ = "STUB: not implemented"; return *new(MultiPolygon) }
