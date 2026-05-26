package geo

import (
	"github.com/paulmach/orb"
)

// Length returns the length of the boundary of the geometry
// using the geo distance function.
func Length(g orb.Geometry) float64 { _ = "STUB: not implemented"; return 0 }

// LengthHaversign returns the length of the boundary of the geometry
// using the geo haversine formula
//
// Deprecated: misspelled, use correctly spelled `LengthHaversine` instead.
func LengthHaversign(g orb.Geometry) float64 { _ = "STUB: not implemented"; return 0 }

// LengthHaversine returns the length of the boundary of the geometry
// using the geo haversine formula
func LengthHaversine(g orb.Geometry) float64 { _ = "STUB: not implemented"; return 0 }
