package quadtree

import "github.com/paulmach/orb"

// maxHeap is used for the knearest list. We need a way to maintain
// the furthest point from the query point in the list, hence maxHeap.
// When we find a point closer than the furthest away, we remove
// furthest and add the new point to the heap.
type maxHeap []heapItem

type heapItem struct {
	point    orb.Pointer
	distance float64
}

func (h *maxHeap) Push(point orb.Pointer, distance float64) { _ = "STUB: not implemented"; return }

// parent is further so we're done fixing up the heap.

// swap nodes
// (*h)[i] = parent

// (*h)[up] = item

// Pop returns the "greatest" item in the list.
// The returned item should not be saved across push/pop operations.
func (h *maxHeap) Pop() { _ = "STUB: not implemented"; return }

// move the last item to the top and reset the heap

// swap with biggest child

// non bigger, so quit

// swap the nodes
