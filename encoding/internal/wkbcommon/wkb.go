package wkbcommon

import (
	"encoding/binary"
	"io"

	"github.com/paulmach/orb"
)

// byteOrder represents little or big endian encoding.
// We don't use binary.ByteOrder because that is an interface
// that leaks to the heap all over the place.
type byteOrder int

const bigEndian byteOrder = 0
const littleEndian byteOrder = 1

const (
	pointType              uint32 = 1
	lineStringType         uint32 = 2
	polygonType            uint32 = 3
	multiPointType         uint32 = 4
	multiLineStringType    uint32 = 5
	multiPolygonType       uint32 = 6
	geometryCollectionType uint32 = 7

	ewkbType uint32 = 0x20000000
)

const (
	// limits so that bad data can't come in and preallocate tons of memory.
	// Well formed data with fewer elements will allocate the correct amount just fine.
	MaxPointsAlloc = 10000
	MaxMultiAlloc  = 100
)

// DefaultByteOrder is the order used for marshalling or encoding
// is none is specified.
var DefaultByteOrder binary.ByteOrder = binary.LittleEndian

// An Encoder will encode a geometry as (E)WKB to the writer given at
// creation time.
type Encoder struct {
	buf []byte

	w     io.Writer
	order binary.ByteOrder
}

// MustMarshal will encode the geometry and panic on error.
// Currently there is no reason to error during geometry marshalling.
func MustMarshal(geom orb.Geometry, srid int, byteOrder ...binary.ByteOrder) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Marshal encodes the geometry with the given byte order.
func Marshal(geom orb.Geometry, srid int, byteOrder ...binary.ByteOrder) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEncoder creates a new Encoder for the given writer.
func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

// SetByteOrder will override the default byte order set when
// the encoder was created.
func (e *Encoder) SetByteOrder(bo binary.ByteOrder) {
	_ = "STUB: not implemented"

	// Encode will write the geometry encoded as (E)WKB to the given writer.
	return
}

func (e *Encoder) Encode(geom orb.Geometry, srid int) error { _ = "STUB: not implemented"; return nil }

// nil values should not write any data. Empty sizes will still
// write an empty version of that type.

// deal with types that are not supported by wkb

func (e *Encoder) writeTypePrefix(t uint32, l int, srid int) error {
	_ = "STUB: not implemented"
	return nil
}

// Decoder can decoder (E)WKB geometry off of the stream.
type Decoder struct {
	r io.Reader
}

// Unmarshal will decode the type into a Geometry.
func Unmarshal(data []byte) (orb.Geometry, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), 0, nil
}

// NewDecoder will create a new (E)WKB decoder.
func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// Decode will decode the next geometry off of the stream.
func (d *Decoder) Decode() (orb.Geometry, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), 0, nil
}

func readByteOrderType(r io.Reader, buf []byte) (byteOrder, uint32, int, error) {
	_ = "STUB: not implemented"
	// the byte order is the first byte
	return *new(byteOrder), 0, 0, nil
}

// the type which is 4 bytes

func readUint32(r io.Reader, order byteOrder, buf []byte) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unmarshalByteOrderType(buf []byte) (byteOrder, uint32, int, []byte, error) {
	_ = "STUB: not implemented"
	return *new(byteOrder), 0, 0, nil, nil
}

// regular wkb, no srid

func byteOrderType(buf []byte) (byteOrder, uint32, error) {
	_ = "STUB: not implemented"
	return *new(byteOrder), 0, nil
}

// the type which is 4 bytes

func unmarshalUint32(order byteOrder, buf []byte) uint32 { _ = "STUB: not implemented"; return 0 }

// GeomLength helps to do preallocation during a marshal.
func GeomLength(geom orb.Geometry, ewkb bool) int { _ = "STUB: not implemented"; return 0 }
