// Package wkb is for decoding ESRI's Well Known Binary (WKB) format
// specification at https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry#Well-known_binary
package wkb

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/internal/wkbcommon"
)

var (
	// ErrUnsupportedDataType is returned by Scan methods when asked to scan
	// non []byte data from the database. This should never happen
	// if the driver is acting appropriately.
	ErrUnsupportedDataType = errors.New("wkb: scan value must be []byte")

	// ErrNotWKB is returned when unmarshalling WKB and the data is not valid.
	ErrNotWKB = errors.New("wkb: invalid data")

	// ErrIncorrectGeometry is returned when unmarshalling WKB data into the wrong type.
	// For example, unmarshaling linestring data into a point.
	ErrIncorrectGeometry = errors.New("wkb: incorrect geometry")

	// ErrUnsupportedGeometry is returned when geometry type is not supported by this lib.
	ErrUnsupportedGeometry = errors.New("wkb: unsupported geometry")
)

var commonErrorMap = map[error]error{
	wkbcommon.ErrUnsupportedDataType: ErrUnsupportedDataType,
	wkbcommon.ErrNotWKB:              ErrNotWKB,
	wkbcommon.ErrNotWKBHeader:        ErrNotWKB,
	wkbcommon.ErrIncorrectGeometry:   ErrIncorrectGeometry,
	wkbcommon.ErrUnsupportedGeometry: ErrUnsupportedGeometry,
}

func mapCommonError(err error) error { _ = "STUB: not implemented"; return nil }

// DefaultByteOrder is the order used for marshalling or encoding
// is none is specified.
var DefaultByteOrder binary.ByteOrder = binary.LittleEndian

// An Encoder will encode a geometry as WKB to the writer given at
// creation time.
type Encoder struct {
	e *wkbcommon.Encoder
}

// MustMarshal will encode the geometry and panic on error.
// Currently there is no reason to error during geometry marshalling.
func MustMarshal(geom orb.Geometry, byteOrder ...binary.ByteOrder) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Marshal encodes the geometry with the given byte order.
func Marshal(geom orb.Geometry, byteOrder ...binary.ByteOrder) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalToHex will encode the geometry into a hex string representation of the binary wkb.
func MarshalToHex(geom orb.Geometry, byteOrder ...binary.ByteOrder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustMarshalToHex will encode the geometry and panic on error.
// Currently there is no reason to error during geometry marshalling.
func MustMarshalToHex(geom orb.Geometry, byteOrder ...binary.ByteOrder) string {
	_ = "STUB: not implemented"
	return ""
}

// NewEncoder creates a new Encoder for the given writer.
func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

// SetByteOrder will override the default byte order set when
// the encoder was created.
func (e *Encoder) SetByteOrder(bo binary.ByteOrder) *Encoder { _ = "STUB: not implemented"; return nil }

// Encode will write the geometry encoded as WKB to the given writer.
func (e *Encoder) Encode(geom orb.Geometry) error { _ = "STUB: not implemented"; return nil }

// Decoder can decoder WKB geometry off of the stream.
type Decoder struct {
	d *wkbcommon.Decoder
}

// Unmarshal will decode the type into a Geometry.
func Unmarshal(data []byte) (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}

// NewDecoder will create a new WKB decoder.
func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// Decode will decode the next geometry off of the stream.
func (d *Decoder) Decode() (orb.Geometry, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), nil
}
