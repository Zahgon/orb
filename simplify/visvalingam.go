package simplify

import (
	"github.com/paulmach/orb"
)

var _ orb.Simplifier = &VisvalingamSimplifier{}

// A VisvalingamSimplifier is a reducer that
// performs the visvalingham algorithm.
type VisvalingamSimplifier struct {
	Threshold float64

	// If 0 defaults to 2 for line, 3 for non-closed rings and 4 for closed rings.
	// The intent is to maintain valid geometry after simplification, however it
	// is still possible for the simplification to create self-intersections.
	ToKeep int
}

// Visvalingam creates a new VisvalingamSimplifier.
// If minPointsToKeep is 0 the algorithm will keep at least 2 points for lines,
// 3 for non-closed rings and 4 for closed rings. However it is still possible
// for the simplification to create self-intersections.
func Visvalingam(threshold float64, minPointsToKeep int) *VisvalingamSimplifier {
	_ = "STUB: not implemented"
	return nil
}

// VisvalingamThreshold runs the Visvalingam-Whyatt algorithm removing
// triangles whose area is below the threshold.
// Will keep at least 2 points for lines, 3 for non-closed rings and 4 for closed rings.
// The intent is to maintain valid geometry after simplification, however it
// is still possible for the simplification to create self-intersections.
func VisvalingamThreshold(threshold float64) *VisvalingamSimplifier {
	_ = "STUB: not implemented"
	return nil
}

// VisvalingamKeep runs the Visvalingam-Whyatt algorithm removing
// triangles of minimum area until we're down to `minPointsToKeep` number of points.
// If minPointsToKeep is 0 the algorithm will keep at least 2 points for lines,
// 3 for non-closed rings and 4 for closed rings. However it is still possible
// for the simplification to create self-intersections.
func VisvalingamKeep(minPointsToKeep int) *VisvalingamSimplifier {
	_ = "STUB: not implemented"
	return nil
}

func (s *VisvalingamSimplifier) simplify(ls orb.LineString, area, wim bool) (orb.LineString, []int) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

// create identify map

// edge cases checked, get on with it
// triangle area is doubled to save the multiply :)

// build the initial minheap linked list.

// internal path items

// final item

// run through the reduction process

// remove current element from linked list

// figure out the new areas

// Stuff to create the priority queue, or min heap.
// Rewriting it here, vs using the std lib, resulted in a 50% performance bump!
type minHeap []*visItem

type visItem struct {
	area       float64 // triangle area
	pointIndex int     // index of point in original path

	// to keep a virtual linked list to help rebuild the triangle areas as we remove points.
	next     *visItem
	previous *visItem

	index int // internal index in heap, for removal and update
}

func (h *minHeap) Push(item *visItem) { _ = "STUB: not implemented"; return }

func (h *minHeap) Pop() *visItem { _ = "STUB: not implemented"; return nil }

func (h minHeap) Update(item *visItem, area float64) { _ = "STUB: not implemented"; return }

// area got smaller

// area got larger

func (h minHeap) up(i int) { _ = "STUB: not implemented"; return }

// parent is smaller so we're done fixing up the heap.

// swap nodes

func (h minHeap) down(i int) { _ = "STUB: not implemented"; return }

// swap with smallest child

// non smaller, so quit

// swap the nodes

func doubleTriangleArea(ls orb.LineString, i1, i2, i3 int) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Simplify will run the simplification for any geometry type.
func (s *VisvalingamSimplifier) Simplify(g orb.Geometry) orb.Geometry {
	_ = "STUB: not implemented"
	return *

	// LineString will simplify the linestring using this simplifier.
	new(orb.Geometry)
}

func (s *VisvalingamSimplifier) LineString(ls orb.LineString) orb.LineString {
	_ = "STUB: not implemented"
	return *

	// MultiLineString will simplify the multi-linestring using this simplifier.
	new(orb.LineString)
}

func (s *VisvalingamSimplifier) MultiLineString(mls orb.MultiLineString) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

// Ring will simplify the ring using this simplifier.
func (s *VisvalingamSimplifier) Ring(r orb.Ring) orb.Ring {
	_ = "STUB: not implemented"

	// Polygon will simplify the polygon using this simplifier.
	return *new(orb.Ring)
}

func (s *VisvalingamSimplifier) Polygon(p orb.Polygon) orb.Polygon {
	_ = "STUB: not implemented"
	return *

	// MultiPolygon will simplify the multi-polygon using this simplifier.
	new(orb.Polygon)
}

func (s *VisvalingamSimplifier) MultiPolygon(mp orb.MultiPolygon) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

// Collection will simplify the collection using this simplifier.
func (s *VisvalingamSimplifier) Collection(c orb.Collection) orb.Collection {
	_ = "STUB: not implemented"
	return *new(orb.Collection)
}
