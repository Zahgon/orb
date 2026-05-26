// Package resample has a couple functions for resampling line geometry
// into more or less evenly spaces points.
package resample

import (
	"github.com/paulmach/orb"
)

// Resample converts the line string into totalPoints-1 evenly spaced segments.
// This function will modify the linestring input.
func Resample(ls orb.LineString, df orb.DistanceFunc, totalPoints int) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}

// precomputes the total distance and intermediate distances

// ToInterval coverts the line string into evenly spaced points of
// about the given distance.
// This function will modify the linestring input.
func ToInterval(ls orb.LineString, df orb.DistanceFunc, dist float64) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}

// precomputes the total distance and intermediate distances

func resample(ls orb.LineString, dists []float64, totalDistance float64, totalPoints int) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}

// start stays the same

// declare here and update had nice performance benefits need to retest

// need to add a point

// move to the next distance we want

// weird round off error on my machine

// past the current point in the original segment, so move to the next one

// end stays the same, to handle round off errors
// for 1, we want the first point

// resampleEdgeCases is used to handle edge case for
// resampling like not enough points and the line string is all the same point.
// will return nil if there are no edge cases. If return true if
// one of these edge cases was found and handled.
func resampleEdgeCases(ls orb.LineString, totalPoints int) (orb.LineString, bool) {
	_ = "STUB: not implemented"
	// degenerate case
	return *new(orb.LineString), false
}

// if all the points are the same, treat as special case.

// extend to be requested length

// contract to be requested length

// precomputeDistances precomputes the total distance and intermediate distances.
func precomputeDistances(ls orb.LineString, df orb.DistanceFunc) (float64, []float64) {
	_ = "STUB: not implemented"
	return 0, nil
}
