package wkbcommon

import (
	"io"

	"github.com/paulmach/orb"
)

func unmarshalPolygon(order byteOrder, data []byte) (orb.Polygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.Polygon), nil
}

// invalid data can come in here and allocate tons of memory.

func readPolygon(r io.Reader, order byteOrder, buf []byte) (orb.Polygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.Polygon), nil
}

// invalid data can come in here and allocate tons of memory.

func (e *Encoder) writePolygon(p orb.Polygon, srid int) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalMultiPolygon(order byteOrder, data []byte) (orb.MultiPolygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon), nil
}

// invalid data can come in here and allocate tons of memory.

func readMultiPolygon(r io.Reader, order byteOrder, buf []byte) (orb.MultiPolygon, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon), nil
}

// invalid data can come in here and allocate tons of memory.

func (e *Encoder) writeMultiPolygon(mp orb.MultiPolygon, srid int) error {
	_ = "STUB: not implemented"
	return nil
}
