package geo

import (
	"github.com/paulmach/orb"
)

// Distance returns the distance between two points on the earth.
func Distance(p1, p2 orb.Point) float64 { _ = "STUB: not implemented"; return 0 }

// fast way using pythagorean theorem on an equirectangular projection

// DistanceHaversine computes the distance on the earth using the
// more accurate haversine formula.
func DistanceHaversine(p1, p2 orb.Point) float64 { _ = "STUB: not implemented"; return 0 }

// Bearing computes the direction one must start traveling on earth
// to be heading from, to the given points.
func Bearing(from, to orb.Point) float64 { _ = "STUB: not implemented"; return 0 }

// Midpoint returns the half-way point along a great circle path between the two points.
func Midpoint(p, p2 orb.Point) orb.Point { _ = "STUB: not implemented"; return *new(orb.Point) }

// convert back to degrees

// PointAtBearingAndDistance returns the point at the given bearing and distance in meters from the point
func PointAtBearingAndDistance(p orb.Point, bearing, distance float64) orb.Point {
	_ = "STUB: not implemented"
	return *new(orb.Point)
}

func PointAtDistanceAlongLine(ls orb.LineString, distance float64) (orb.Point, float64) {
	_ = "STUB: not implemented"
	return *new(orb.Point), 0
}
