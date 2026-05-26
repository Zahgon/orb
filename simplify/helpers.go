// Package simplify implements several reducing/simplifying functions for `orb.Geometry` types.
package simplify

import "github.com/paulmach/orb"

type simplifier interface {
	simplify(l orb.LineString, area bool, withIndexMap bool) (orb.LineString, []int)
}

func simplify(s simplifier, geom orb.Geometry) orb.Geometry {
	_ = "STUB: not implemented"
	return *new(orb.Geometry)
}

func lineString(s simplifier, ls orb.LineString) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}

func multiLineString(s simplifier, mls orb.MultiLineString) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

func ring(s simplifier, r orb.Ring) orb.Ring { _ = "STUB: not implemented"; return *new(orb.Ring) }

func polygon(s simplifier, p orb.Polygon) orb.Polygon {
	_ = "STUB: not implemented"
	return *new(orb.Polygon)
}

func multiPolygon(s simplifier, mp orb.MultiPolygon) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

func collection(s simplifier, c orb.Collection) orb.Collection {
	_ = "STUB: not implemented"
	return *new(orb.Collection)
}

func runSimplify(s simplifier, ls orb.LineString, area bool) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}
