package wkbcommon

import (
	"io"

	"github.com/paulmach/orb"
)

func readCollection(r io.Reader, order byteOrder, buf []byte) (orb.Collection, error) {
	_ = "STUB: not implemented"
	return *new(orb.Collection), nil
}

// invalid data can come in here and allocate tons of memory.

func (e *Encoder) writeCollection(c orb.Collection, srid int) error {
	_ = "STUB: not implemented"
	return nil
}
