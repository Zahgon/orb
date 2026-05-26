package wkbcommon

import (
	"io"

	"github.com/paulmach/orb"
)

func unmarshalPoints(order byteOrder, data []byte) ([]orb.Point, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// invalid data can come in here and allocate tons of memory.

func unmarshalPoint(order byteOrder, buf []byte) (orb.Point, error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), nil
}

func readPoint(r io.Reader, order byteOrder, buf []byte) (orb.Point, error) {
	_ = "STUB: not implemented"
	return *new(orb.Point), nil
}

func (e *Encoder) writePoint(p orb.Point, srid int) error { _ = "STUB: not implemented"; return nil }

func unmarshalMultiPoint(order byteOrder, data []byte) (orb.MultiPoint, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint), nil
}

// invalid data can come in here and allocate tons of memory.

func readMultiPoint(r io.Reader, order byteOrder, buf []byte) (orb.MultiPoint, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiPoint), nil
}

// invalid data can come in here and allocate tons of memory.

func (e *Encoder) writeMultiPoint(mp orb.MultiPoint, srid int) error {
	_ = "STUB: not implemented"
	return nil
}
