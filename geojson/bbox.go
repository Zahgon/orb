package geojson

import "github.com/paulmach/orb"

// BBox is for the geojson bbox attribute which is an array with all axes
// of the most southwesterly point followed by all axes of the more northeasterly point.
type BBox []float64

// NewBBox creates a bbox from a a bound.
func NewBBox(b orb.Bound) BBox { _ = "STUB: not implemented"; return *new(BBox) }

// Valid checks if the bbox is present and has at least 4 elements.
func (bb BBox) Valid() bool { _ = "STUB: not implemented"; return false }

// Bound returns the orb.Bound for the BBox.
func (bb BBox) Bound() orb.Bound { _ = "STUB: not implemented"; return *new(orb.Bound) }
