package smartclip

import (
	"github.com/paulmach/orb"
)

// aroundBound will connect the endpoints of the linestring provided
// by wrapping the line around the bounds in the direction provided.
// Will append to the input.
func aroundBound(
	box orb.Bound,
	in orb.Ring,
	o orb.Orientation,
) orb.Ring {
	_ = "STUB: not implemented"
	return *new(orb.Ring)
}

// endpoints long an edge. Need to figure out what order they're in
// to figure out if we just need to connect them or go all the way around.

// move to next and go until we're all the way around.

// add first point to the end to make it a ring

//         left  mid  right
//    top  1001  1000  1010
//    mid  0001  0000  0010
// bottom  0101  0100  0110

// on the boundary is outside
func bitCodeOpen(b orb.Bound, p orb.Point) int { _ = "STUB: not implemented"; return 0 }

// pointFor returns a representative point for the side of the given bitCode.
func pointFor(b orb.Bound, code int) orb.Point { _ = "STUB: not implemented"; return *new(orb.Point) }

//         left  mid  right
//    top     9     8    10
//    mid     1     0     2
// bottom     5     4     6

// nexts takes a bitcode index and jumps to the next corner.
var nexts = map[orb.Orientation][11]int{
	orb.CW: {
		-1,
		9, // 1
		6, // 2
		-1,
		5, // 4
		1, // 5
		4, // 6
		-1,
		10, // 8
		8,  // 9
		2,  // 10
	},
	orb.CCW: {
		-1,
		5,  // 1
		10, // 2
		-1,
		6, // 4
		4, // 5
		2, // 6
		-1,
		9, // 8
		1, // 9
		8, // 10
	},
}
