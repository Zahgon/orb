package tilecover

import "github.com/paulmach/orb/maptile"

// MergeUp will merge up the tiles in a given set up to the
// given min zoom. Tiles will be merged up only if all 4 siblings
// are in the set. The tiles in the input set are expected
// to all be of the same zoom, e.g. outputs of the Geometry function.
func MergeUp(set maptile.Set, min maptile.Zoom) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}

// MergeUpPartial will merge up the tiles in a given set up to the
// given min zoom. Tiles will be merged up if `count` siblings are in the
// set. The tiles in the input set are expected to all be of the same
// zoom, e.g. outputs of the Geometry function.
func MergeUpPartial(set maptile.Set, min maptile.Zoom, count int) maptile.Set {
	_ = "STUB: not implemented"
	return *new(maptile.Set)
}
