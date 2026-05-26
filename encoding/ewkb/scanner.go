package ewkb

import (
	"database/sql"
	"database/sql/driver"

	"github.com/paulmach/orb"
)

var (
	_ sql.Scanner  = &GeometryScanner{}
	_ driver.Value = value{}
)

// GeometryScanner is a thing that can scan in sql query results.
// It can be used as a scan destination:
//
//	var s wkb.GeometryScanner
//	err := db.QueryRow("SELECT latlon FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	  // use s.Geometry
//	  // use s.SRID
//	} else {
//	  // NULL value
//	}
type GeometryScanner struct {
	sridInPrefix bool
	g            any
	SRID         int
	Geometry     orb.Geometry
	Valid        bool // Valid is true if the geometry is not NULL
}

// Scanner will return a GeometryScanner that can scan sql query results.
// The geometryScanner.Geometry attribute will be set to the value.
// If g is non-nil, it MUST be a pointer to an orb.Geometry
// type like a Point or LineString. In that case the value will be written to
// g and the Geometry attribute.
//
//	var p orb.Point
//	err := db.QueryRow("SELECT latlon FROM foo WHERE id=?", id).Scan(wkb.Scanner(&p))
//	...
//	// use p
//
// If the value may be null check Valid first:
//
//	var point orb.Point
//	s := wkb.Scanner(&point)
//	err := db.QueryRow("SELECT latlon FROM foo WHERE id=?", id).Scan(s)
//	...
//	if s.Valid {
//	  // use p
//	} else {
//	  // NULL value
//	}
func Scanner(g any) *GeometryScanner { _ = "STUB: not implemented"; return nil }

// ScannerPrefixSRID will scan ewkb data were the SRID is in the first 4 bytes of the data.
// Databases like mysql/mariadb use this as their raw format. This method should only be used when
// working with such a database.
//
//	var p orb.Point
//	err := db.QueryRow("SELECT latlon FROM foo WHERE id=?", id).Scan(wkb.PrefixSRIDScanner(&p))
//
// However, it is recommended to covert to wkb explicitly using something like:
//
//	var srid int
//	var p orb.Point
//	err := db.QueryRow("SELECT ST_SRID(latlon), ST_AsBinary(latlon) FROM foo WHERE id=?", id).
//		Scan(&srid, wkb.Scanner(&p))
//
// https://dev.mysql.com/doc/refman/5.7/en/gis-data-formats.html
func ScannerPrefixSRID(g any) *GeometryScanner { _ = "STUB: not implemented"; return nil }

// Scan will scan the input []byte data into a geometry.
// This could be into the orb geometry type pointer or, if nil,
// the scanner.Geometry attribute.
func (s *GeometryScanner) Scan(d any) error { _ = "STUB: not implemented"; return nil }

type value struct {
	srid int
	v    orb.Geometry
}

// Value will create a driver.Valuer that will EWKB the geometry into the database query.
//
//	db.Exec("INSERT INTO table (point_column) VALUES (?)", ewkb.Value(p, 4326))
func Value(g orb.Geometry, srid int) driver.Valuer {
	_ = "STUB: not implemented"
	return *new(driver.Valuer)
}

func (v value) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

type valuePrefixSRID struct {
	srid int
	v    orb.Geometry
}

// ValuePrefixSRID will create a driver.Valuer that will WKB the geometry
// but add the srid as a 4 byte prefix.
//
//	db.Exec("INSERT INTO table (point_column) VALUES (?)", ewkb.Value(p, 4326))
func ValuePrefixSRID(g orb.Geometry, srid int) driver.Valuer {
	_ = "STUB: not implemented"
	return *new(driver.Valuer)
}

func (v valuePrefixSRID) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
