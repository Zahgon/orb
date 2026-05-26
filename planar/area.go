// Package planar computes properties on geometries assuming they are
// in 2d euclidean space.
package planar

import (
	"github.com/paulmach/orb"
)

// Area returns the area of the geometry in the 2d plane.
func Area(g orb.Geometry) float64 {
	_ = "STUB: not implemented"
	// TODO: make faster non-centroid version.
	return 0
}

// CentroidArea returns both the centroid and the area in the 2d plane.
// Since the area is need for the centroid, return both.
// Polygon area will always be >= zero. Ring area may be negative if it has
// a clockwise winding order.
func CentroidArea(g orb.Geometry) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}

func multiPointCentroid(mp orb.MultiPoint) orb.Point {
	_ = "STUB: not implemented"
	return *new(orb.Point)
}

func multiLineStringCentroid(mls orb.MultiLineString) orb.Point {
	_ = "STUB: not implemented"
	return *new(orb.Point)
}

func lineStringCentroidDist(ls orb.LineString) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}

// implicitly move everything to near the origin to help with roundoff

func ringCentroidArea(r orb.Ring) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}

// implicitly move everything to near the origin to help with roundoff

// no need to deal with first and last vertex since we "moved"
// that point the origin (multiply by 0 == 0)

func polygonCentroidArea(p orb.Polygon) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}

func multiPolygonCentroidArea(mp orb.MultiPolygon) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}

func collectionCentroidArea(c orb.Collection) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}

func maxDim(c orb.Collection) int { _ = "STUB: not implemented"; return 0 }
