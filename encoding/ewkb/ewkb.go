package ewkb

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

	// ErrNotEWKB is returned when unmarshalling EWKB and the data is not valid.
	ErrNotEWKB = errors.New("wkb: invalid data")

	// ErrIncorrectGeometry is returned when unmarshalling EWKB data into the wrong type.
	// For example, unmarshaling linestring data into a point.
	ErrIncorrectGeometry = errors.New("wkb: incorrect geometry")

	// ErrUnsupportedGeometry is returned when geometry type is not supported by this lib.
	ErrUnsupportedGeometry = errors.New("wkb: unsupported geometry")
)

var commonErrorMap = map[error]error{
	wkbcommon.ErrUnsupportedDataType: ErrUnsupportedDataType,
	wkbcommon.ErrNotWKB:              ErrNotEWKB,
	wkbcommon.ErrNotWKBHeader:        ErrNotEWKB,
	wkbcommon.ErrIncorrectGeometry:   ErrIncorrectGeometry,
	wkbcommon.ErrUnsupportedGeometry: ErrUnsupportedGeometry,
}

func mapCommonError(err error) error { _ = "STUB: not implemented"; return nil }

// DefaultByteOrder is the order used for marshalling or encoding is none is specified.
var DefaultByteOrder binary.ByteOrder = binary.LittleEndian

// DefaultSRID is set to 4326, a common SRID, which represents spatial data using
// longitude and latitude coordinates on the Earth's surface as defined in the WGS84 standard,
// which is also used for the Global Positioning System (GPS).
// This will be used by the encoder if non is specified.
var DefaultSRID int = 4326

// An Encoder will encode a geometry as EWKB to the writer given at creation time.
type Encoder struct {
	srid int
	e    *wkbcommon.Encoder
}

// MustMarshal will encode the geometry and panic on error.
// Currently there is no reason to error during geometry marshalling.
func MustMarshal(geom orb.Geometry, srid int, byteOrder ...binary.ByteOrder) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Marshal encodes the geometry with the given byte order.
// An SRID of 0 will not be included in the encoding and the result will be a wkb encoding of the geometry.
func Marshal(geom orb.Geometry, srid int, byteOrder ...binary.ByteOrder) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalToHex will encode the geometry into a hex string representation of the binary ewkb.
func MarshalToHex(geom orb.Geometry, srid int, byteOrder ...binary.ByteOrder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MustMarshalToHex will encode the geometry and panic on error.
// Currently there is no reason to error during geometry marshalling.
func MustMarshalToHex(geom orb.Geometry, srid int, byteOrder ...binary.ByteOrder) string {
	_ = "STUB: not implemented"
	return ""
}

// NewEncoder creates a new Encoder for the given writer.
func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

// SetByteOrder will override the default byte order set when
// the encoder was created.
func (e *Encoder) SetByteOrder(bo binary.ByteOrder) *Encoder { _ = "STUB: not implemented"; return nil }

// SetSRID will override the default srid.
func (e *Encoder) SetSRID(srid int) *Encoder { _ = "STUB: not implemented"; return nil }

// Encode will write the geometry encoded as EWKB to the given writer.
func (e *Encoder) Encode(geom orb.Geometry, srid ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// Decoder can decoder WKB geometry off of the stream.
type Decoder struct {
	d *wkbcommon.Decoder
}

// Unmarshal will decode the type into a Geometry.
func Unmarshal(data []byte) (orb.Geometry, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), 0, nil
}

// NewDecoder will create a new EWKB decoder.
func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// Decode will decode the next geometry off of the stream.
func (d *Decoder) Decode() (orb.Geometry, int, error) {
	_ = "STUB: not implemented"
	return *new(orb.Geometry), 0, nil
}
