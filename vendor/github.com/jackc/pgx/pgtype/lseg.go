package pgtype

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/jackc/pgx/pgio"
)

type Lseg struct {
	P      [2]Vec2
	Status Status
}

func (dst *Lseg) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to Lseg", src)
}

func (dst *Lseg) Get() interface{} {
	switch dst.Status {
	case Present:
		return dst
	case Null:
		return nil
	default:
		return dst.Status
	}
}

func (src *Lseg) AssignTo(dst interface{}) error {
	return fmt.Errorf("cannot assign %v to %T", src, dst)
}

func (dst *Lseg) DecodeText(ci *ConnInfo, src []byte) error {
	if src == nil {
		*dst = Lseg{Status: Null}
		return nil
	}

	if len(src) < 11 {
		return fmt.Errorf("invalid length for Lseg: %v", len(src))
	}

	str := string(src[2:])

	var end int
	end = strings.IndexByte(str, ',')

	x1, err := strconv.ParseFloat(str[:end], 64)
	if err != nil {
		return err
	}

	str = str[end+1:]
	end = strings.IndexByte(str, ')')

	y1, err := strconv.ParseFloat(str[:end], 64)
	if err != nil {
		return err
	}

	str = str[end+3:]
	end = strings.IndexByte(str, ',')

	x2, err := strconv.ParseFloat(str[:end], 64)
	if err != nil {
		return err
	}

	str = str[end+1 : len(str)-2]

	y2, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*dst = Lseg{P: [2]Vec2{{x1, y1}, {x2, y2}}, Status: Present}
	return nil
}

func (dst *Lseg) DecodeBinary(ci *ConnInfo, src []byte) error {
	if src == nil {
		*dst = Lseg{Status: Null}
		return nil
	}

	if len(src) != 32 {
		return fmt.Errorf("invalid length for Lseg: %v", len(src))
	}

	x1 := binary.BigEndian.Uint64(src)
	y1 := binary.BigEndian.Uint64(src[8:])
	x2 := binary.BigEndian.Uint64(src[16:])
	y2 := binary.BigEndian.Uint64(src[24:])

	*dst = Lseg{
		P: [2]Vec2{
			{math.Float64frombits(x1), math.Float64frombits(y1)},
			{math.Float64frombits(x2), math.Float64frombits(y2)},
		},
		Status: Present,
	}
	return nil
}

func (src *Lseg) EncodeText(ci *ConnInfo, w io.Writer) (bool, error) {
	switch src.Status {
	case Null:
		return true, nil
	case Undefined:
		return false, errUndefined
	}

	_, err := io.WriteString(w, fmt.Sprintf(`(%f,%f),(%f,%f)`,
		src.P[0].X, src.P[0].Y, src.P[1].X, src.P[1].Y))
	return false, err
}

func (src *Lseg) EncodeBinary(ci *ConnInfo, w io.Writer) (bool, error) {
	switch src.Status {
	case Null:
		return true, nil
	case Undefined:
		return false, errUndefined
	}

	if _, err := pgio.WriteUint64(w, math.Float64bits(src.P[0].X)); err != nil {
		return false, err
	}

	if _, err := pgio.WriteUint64(w, math.Float64bits(src.P[0].Y)); err != nil {
		return false, err
	}

	if _, err := pgio.WriteUint64(w, math.Float64bits(src.P[1].X)); err != nil {
		return false, err
	}

	_, err := pgio.WriteUint64(w, math.Float64bits(src.P[1].Y))
	return false, err
}

// Scan implements the database/sql Scanner interface.
func (dst *Lseg) Scan(src interface{}) error {
	if src == nil {
		*dst = Lseg{Status: Null}
		return nil
	}

	switch src := src.(type) {
	case string:
		return dst.DecodeText(nil, []byte(src))
	case []byte:
		return dst.DecodeText(nil, src)
	}

	return fmt.Errorf("cannot scan %T", src)
}

// Value implements the database/sql/driver Valuer interface.
func (src *Lseg) Value() (driver.Value, error) {
	return encodeValueText(src)
}
