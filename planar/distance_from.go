package planar

import (
	"github.com/paulmach/orb"
)

// DistanceFromSegment returns the point's distance from the segment [a, b].
func DistanceFromSegment(a, b, point orb.Point) float64 { _ = "STUB: not implemented"; return 0 }

// DistanceFromSegmentSquared returns point's squared distance from the segment [a, b].
func DistanceFromSegmentSquared(a, b, point orb.Point) float64 { _ = "STUB: not implemented"; return 0 }

// DistanceFrom returns the distance from the boundary of the geometry in
// the units of the geometry.
func DistanceFrom(g orb.Geometry, p orb.Point) float64 { _ = "STUB: not implemented"; return 0 }

// DistanceFromWithIndex returns the minimum euclidean distance
// from the boundary of the geometry plus the index of the sub-geometry
// that was the match.
func DistanceFromWithIndex(g orb.Geometry, p orb.Point) (float64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func multiPointDistanceFrom(mp orb.MultiPoint, p orb.Point) (float64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func lineStringDistanceFrom(ls orb.LineString, p orb.Point) (float64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func polygonDistanceFrom(p orb.Polygon, point orb.Point) (float64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func segmentDistanceFromSquared(p1, p2, point orb.Point) float64 {
	_ = "STUB: not implemented"
	return 0
}
