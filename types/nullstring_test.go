package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strconv"
	"testing"
)

var (
	stringJSON      = []byte(`"test"`)
	blankStringJSON = []byte(`""`)
	nullStringJSON  = []byte(`{"String":"test","Valid":true}`)

	nullJSON          = []byte(`null`)
	invalidStringJSON = []byte(`:)`)
)

func TestMarshalZeroString(t *testing.T) {
	v1, _ := json.Marshal("")

	x := NullString{sql.NullString{String: "", Valid: true}}
	v2, _ := x.MarshalJSON()

	// Equal to 0 if two byte slices match.
	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNullString(t *testing.T) {
	v1, _ := json.Marshal(nil)

	x := NullString{sql.NullString{String: "", Valid: false}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalString(t *testing.T) {
	v1, _ := json.Marshal("`")

	x := NullString{sql.NullString{String: "`", Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestUnmarshalZeroString(t *testing.T) {
	v1 := NullString{sql.NullString{String: "", Valid: false}}
	v2 := NullString{sql.NullString{}}

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
	if v1.String != v2.String || v1.Valid != v2.Valid {
		t.Logf("%v and %v does not match!", v1, v2)
		t.Fail()
	}
}

func TestUnmarshalString(t *testing.T) {
	var str NullString
	err := str.UnmarshalJSON(stringJSON)

	var ns NullString
	err = ns.UnmarshalJSON(nullStringJSON)
	if err == nil {
	}

	var blank NullString
	err = blank.UnmarshalJSON(blankStringJSON)
	if blank.Valid == false {
		t.Error("blank string should be valid")
	}

	var null NullString
	err = null.UnmarshalJSON(nullJSON)

	var badType NullString
	err = badType.UnmarshalJSON(boolJSON)
	if err == nil {
	}

}

func TestTextUnmarshalString(t *testing.T) {
	var str NullString
	err := str.UnmarshalJSON([]byte(""))
	if err == nil {
		t.Fail()
	}

	var null NullString
	err = null.UnmarshalJSON([]byte(""))
}
