package orb

// Geometry is an interface that represents the shared attributes
// of a geometry.
type Geometry interface {
	GeoJSONType() string
	Dimensions() int // e.g. 0d, 1d, 2d
	Bound() Bound

	// requiring because sub package type switch over all possible types.
	private()
}

// compile time checks
var (
	_ Geometry = Point{}
	_ Geometry = MultiPoint{}
	_ Geometry = LineString{}
	_ Geometry = MultiLineString{}
	_ Geometry = Ring{}
	_ Geometry = Polygon{}
	_ Geometry = MultiPolygon{}
	_ Geometry = Bound{}

	_ Geometry = Collection{}
)

func (p Point) private()             { _ = "STUB: not implemented"; return }
func (mp MultiPoint) private()       { _ = "STUB: not implemented"; return }
func (ls LineString) private()       { _ = "STUB: not implemented"; return }
func (mls MultiLineString) private() { _ = "STUB: not implemented"; return }
func (r Ring) private()              { _ = "STUB: not implemented"; return }
func (p Polygon) private()           { _ = "STUB: not implemented"; return }
func (mp MultiPolygon) private()     { _ = "STUB: not implemented"; return }
func (b Bound) private()             { _ = "STUB: not implemented"; return }
func (c Collection) private() {
	_ = "STUB: not implemented"

	// AllGeometries lists all possible types and values that a geometry
	// interface can be. It should be used only for testing to verify
	// functions that accept a Geometry will work in all cases.
	return
}

var AllGeometries = []Geometry{
	nil,
	Point{},
	MultiPoint{},
	LineString{},
	MultiLineString{},
	Ring{},
	Polygon{},
	MultiPolygon{},
	Bound{},
	Collection{},

	// nil values
	MultiPoint(nil),
	LineString(nil),
	MultiLineString(nil),
	Ring(nil),
	Polygon(nil),
	MultiPolygon(nil),
	Collection(nil),

	// Collection of Collection
	Collection{Collection{Point{}}},
}

// A Collection is a collection of geometries that is also a Geometry.
type Collection []Geometry

// GeoJSONType returns the geometry collection type.
func (c Collection) GeoJSONType() string { _ = "STUB: not implemented"; return "" }

// Dimensions returns the max of the dimensions of the collection.
func (c Collection) Dimensions() int { _ = "STUB: not implemented"; return 0 }

// Bound returns the bounding box of all the Geometries combined.
func (c Collection) Bound() Bound { _ = "STUB: not implemented"; return *new(Bound) }

// Equal compares two collections. Returns true if lengths are the same
// and all the sub geometries are the same and in the same order.
func (c Collection) Equal(collection Collection) bool { _ = "STUB: not implemented"; return false }

// Clone returns a deep copy of the collection.
func (c Collection) Clone() Collection { _ = "STUB: not implemented"; return *new(Collection) }
