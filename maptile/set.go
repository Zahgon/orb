package maptile

import (
	"github.com/paulmach/orb/geojson"
)

// Set is a map/hash of tiles.
type Set map[Tile]bool

// ToFeatureCollection converts a set of tiles into a feature collection.
// This method is mostly useful for debugging output.
func (s Set) ToFeatureCollection() *geojson.FeatureCollection {
	_ = "STUB: not implemented"
	return nil
}

// Merge will merge the given set into the existing set.
func (s Set) Merge(set Set) { _ = "STUB: not implemented"; return }
