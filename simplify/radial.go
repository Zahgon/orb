package simplify

import (
	"github.com/paulmach/orb"
)

var _ orb.Simplifier = &RadialSimplifier{}

// A RadialSimplifier wraps the Radial functions
type RadialSimplifier struct {
	DistanceFunc orb.DistanceFunc
	Threshold    float64 // euclidean distance
}

// Radial creates a new RadialSimplifier.
func Radial(df orb.DistanceFunc, threshold float64) *RadialSimplifier {
	_ = "STUB: not implemented"
	return nil
}

func (s *RadialSimplifier) simplify(ls orb.LineString, area, wim bool) (orb.LineString, []int) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

// Simplify will run the simplification for any geometry type.
func (s *RadialSimplifier) Simplify(g orb.Geometry) orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// LineString will simplify the linestring using this simplifier.
	new(orb.Geometry)
}

func (s *RadialSimplifier) LineString(ls orb.LineString) orb.LineString {
	_ = "STUB: not implemented"
	return *

	// MultiLineString will simplify the multi-linestring using this simplifier.
	new(orb.LineString)
}

func (s *RadialSimplifier) MultiLineString(mls orb.MultiLineString) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// Ring will simplify the ring using this simplifier.
func (s *RadialSimplifier) Ring(r orb.Ring) orb.Ring {
	_ = "STUB: not implemented"

	// Polygon will simplify the polygon using this simplifier.
	return *new(orb.Ring)
}

func (s *RadialSimplifier) Polygon(p orb.Polygon) orb.Polygon {
	_ = "STUB: not implemented"
	return *

	// MultiPolygon will simplify the multi-polygon using this simplifier.
	new(orb.Polygon)
}

func (s *RadialSimplifier) MultiPolygon(mp orb.MultiPolygon) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// Collection will simplify the collection using this simplifier.
func (s *RadialSimplifier) Collection(c orb.Collection) orb.Collection {
	_ = "STUB: not implemented"
	return *new(orb.Collection)
}
