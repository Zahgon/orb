package wkbcommon

import (
	"io"

	"github.com/paulmach/orb"
)

func unmarshalLineString(order byteOrder, data []byte) (orb.LineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

func readLineString(r io.Reader, order byteOrder, buf []byte) (orb.LineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.LineString), nil
}

// invalid data can come in here and allocate tons of memory.

func (e *Encoder) writeLineString(ls orb.LineString, srid int) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalMultiLineString(order byteOrder, data []byte) (orb.MultiLineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString), nil
}

// invalid data can come in here and allocate tons of memory.

func readMultiLineString(r io.Reader, order byteOrder, buf []byte) (orb.MultiLineString, error) {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString), nil
}

// invalid data can come in here and allocate tons of memory.

func (e *Encoder) writeMultiLineString(mls orb.MultiLineString, srid int) error {
	_ = "STUB: not implemented"
	return nil
}
