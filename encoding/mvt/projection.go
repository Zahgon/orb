package mvt

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/maptile"
)

type projection struct {
	ToTile  orb.Projection
	ToWGS84 orb.Projection
}

func newProjection(tile maptile.Tile, extent uint32) *projection {
	_ = "STUB: not implemented"
	return nil

	// powers of two extents allows for some more simplicity
}

func nonPowerOfTwoProjection(tile maptile.Tile, extent uint32) *projection {
	_ = "STUB: not implemented"
	// I really don't know why anyone would use a non-power of two extent,
	// but technically it is supported.
	return nil
}

func isPowerOfTwo(n uint32) bool { _ = "STUB: not implemented"; return false }
