package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"math"
	"strconv"
	"testing"
)

func TestMarshalZeroInt32(t *testing.T) {
	v1, _ := json.Marshal(0)

	x := NullInt32{sql.NullInt32{Int32: 0, Valid: true}}
	v2, _ := x.MarshalJSON()

	// Equal to 0 if two byte slices match.
	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNulInt32(t *testing.T) {
	v1, _ := json.Marshal(nil)

	x := NullInt32{sql.NullInt32{Int32: 0, Valid: false}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalPositiveInt32(t *testing.T) {
	v1, _ := json.Marshal(20000)

	x := NullInt32{sql.NullInt32{Int32: 20000, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNegativeInt32(t *testing.T) {
	v1, _ := json.Marshal(-31492)

	x := NullInt32{sql.NullInt32{Int32: -31492, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestUnmarshalZeroInt32(t *testing.T) {
	v1 := NullInt32{sql.NullInt32{Int32: 0, Valid: true}}
	v2 := NullInt32{}

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
	if v1.Int32 != v2.Int32 || v1.Valid != v2.Valid {
		t.Logf("%v and %v does not match!", v1, v2)
		t.Fail()
	}
}

func TestUnmarshalInt32(t *testing.T) {
	var i NullInt32
	err := i.UnmarshalJSON(invalidIntJSON)

	var si NullInt32
	err = si.UnmarshalJSON(intStringJSON)

	var ni NullInt32
	err = ni.UnmarshalJSON(nullIntJSON)
	if err == nil {
		panic("err should not be nill")
	}

	var null NullInt32
	err = null.UnmarshalJSON(nullJSON)

	var badType NullInt32
	err = badType.UnmarshalJSON(boolJSON)
	if err == nil {
		panic("err should not be nil")
	}

}

func TestUnmarshalInt32Overflow(t *testing.T) {
	int64Overflow := uint64(math.MaxInt64)

	// Max int64 should decode successfully
	var i NullInt32
	err := i.UnmarshalJSON([]byte(strconv.FormatUint(int64Overflow, 10)))

	// Attempt to overflow
	int64Overflow++
	err = i.UnmarshalJSON([]byte(strconv.FormatUint(int64Overflow, 10)))
	if err == nil {
		panic("err should be present; decoded value overflows int64")
	}
}

func TestTextUnmarshalInt32(t *testing.T) {
	var i NullInt32
	err := i.UnmarshalJSON([]byte("12345"))

	var blank NullInt32
	err = blank.UnmarshalJSON([]byte(""))

	var null NullInt32
	err = null.UnmarshalJSON([]byte("null"))

	var invalid NullInt32
	err = invalid.UnmarshalJSON([]byte("hello world"))
	if err == nil {
		panic("expected error")
	}
}
