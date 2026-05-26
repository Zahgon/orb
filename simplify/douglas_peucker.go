package simplify

import (
	"github.com/paulmach/orb"
)

var _ orb.Simplifier = &DouglasPeuckerSimplifier{}

// A DouglasPeuckerSimplifier wraps the DouglasPeucker function.
type DouglasPeuckerSimplifier struct {
	Threshold float64
}

// DouglasPeucker creates a new DouglasPeuckerSimplifier.
func DouglasPeucker(threshold float64) *DouglasPeuckerSimplifier {
	_ = "STUB: not implemented"
	return nil
}

func (s *DouglasPeuckerSimplifier) simplify(ls orb.LineString, area, wim bool) (orb.LineString, []int) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

// dpWorker does the recursive threshold checks.
// Using a stack array with a stackLength variable resulted in
// 4x speed improvement over calling the function recursively.
func dpWorker(ls orb.LineString, threshold float64, mask []byte) int {
	_ = "STUB: not implemented"
	return 0
}

// modify the line in place

// Simplify will run the simplification for any geometry type.
func (s *DouglasPeuckerSimplifier) Simplify(g orb.Geometry) orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// LineString will simplify the linestring using this simplifier.
	new(orb.Geometry)
}

func (s *DouglasPeuckerSimplifier) LineString(ls orb.LineString) orb.LineString {
	_ = "STUB: not implemented"
	return *

	// MultiLineString will simplify the multi-linestring using this simplifier.
	new(orb.LineString)
}

func (s *DouglasPeuckerSimplifier) MultiLineString(mls orb.MultiLineString) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// Ring will simplify the ring using this simplifier.
func (s *DouglasPeuckerSimplifier) Ring(r orb.Ring) orb.Ring {
	_ = "STUB: not implemented"

	// Polygon will simplify the polygon using this simplifier.
	return *new(orb.Ring)
}

func (s *DouglasPeuckerSimplifier) Polygon(p orb.Polygon) orb.Polygon {
	_ = "STUB: not implemented"
	return *

	// MultiPolygon will simplify the multi-polygon using this simplifier.
	new(orb.Polygon)
}

func (s *DouglasPeuckerSimplifier) MultiPolygon(mp orb.MultiPolygon) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// Collection will simplify the collection using this simplifier.
func (s *DouglasPeuckerSimplifier) Collection(c orb.Collection) orb.Collection {
	_ = "STUB: not implemented"
	return *new(orb.Collection)
}
