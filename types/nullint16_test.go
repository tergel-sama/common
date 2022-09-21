package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"math"
	"strconv"
	"testing"
)

var (
	intJSON        = []byte(`12345`)
	invalidIntJSON = []byte(`12.345`)
	intStringJSON  = []byte(`"12345"`)
	nullIntJSON    = []byte(`{"int16":12345,"Valid":true}`)
)

func TestMarshalZeroInt16(t *testing.T) {
	v1, _ := json.Marshal(0)

	x := NullInt16{sql.NullInt16{Int16: 0, Valid: true}}
	v2, _ := x.MarshalJSON()

	// Equal to 0 if two byte slices match.
	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNulInt16(t *testing.T) {
	v1, _ := json.Marshal(nil)

	x := NullInt16{sql.NullInt16{Int16: 0, Valid: false}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalPositiveInt16(t *testing.T) {
	v1, _ := json.Marshal(20000)

	x := NullInt16{sql.NullInt16{Int16: 20000, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNegativeInt16(t *testing.T) {
	v1, _ := json.Marshal(-31492)

	x := NullInt16{sql.NullInt16{Int16: -31492, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestUnmarshalZeroInt16(t *testing.T) {
	v1 := NullInt16{sql.NullInt16{Int16: 0, Valid: true}}
	v2 := NullInt16{}

	x := []byte(strconv.Itoa(0))

	err := v2.UnmarshalJSON(x)
	if err != nil {
		t.Logf("%v", err)
	}

	// Notes on comparison:
	// reflect.DeepEqual may not return the expected result
	// if underlying fields have different metadata (i.e., timezone difference).

	// cmp.Equal is intended to only be used in tests, as performance is not a goal and
	// it may panic if it cannot compare the values.

	// Or compare the two manually.
	if v1.Int16 != v2.Int16 || v1.Valid != v2.Valid {
		t.Logf("%v and %v does not match!", v1, v2)
		t.Fail()
	}
}

func TestUnmarshalInt16(t *testing.T) {
	var i NullInt16
	err := i.UnmarshalJSON(invalidIntJSON)

	var si NullInt16
	err = si.UnmarshalJSON(intStringJSON)

	var ni NullInt16
	err = ni.UnmarshalJSON(nullIntJSON)
	if err == nil {
		panic("err should not be nill")
	}

	var null NullInt16
	err = null.UnmarshalJSON(nullJSON)

	var badType NullInt16
	err = badType.UnmarshalJSON(boolJSON)
	if err == nil {
		panic("err should not be nil")
	}

}

func TestUnmarshalInt16Overflow(t *testing.T) {
	int16Overflow := uint64(math.MaxInt64)

	// Max int16 should decode successfully
	var i NullInt16
	err := i.UnmarshalJSON([]byte(strconv.FormatUint(int16Overflow, 10)))

	// Attempt to overflow
	int16Overflow++
	err = i.UnmarshalJSON([]byte(strconv.FormatUint(int16Overflow, 10)))
	if err == nil {
		panic("err should be present; decoded value overflows int16")
	}
}

func TestTextUnmarshalInt16(t *testing.T) {
	var i NullInt16
	err := i.UnmarshalJSON([]byte("12345"))

	var blank NullInt16
	err = blank.UnmarshalJSON([]byte(""))

	var null NullInt16
	err = null.UnmarshalJSON([]byte("null"))

	var invalid NullInt16
	err = invalid.UnmarshalJSON([]byte("hello world"))
	if err == nil {
		panic("expected error")
	}
}
