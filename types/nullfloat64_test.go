package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strconv"
	"testing"
)

var (
	floatJSON       = []byte(`1.2345`)
	floatStringJSON = []byte(`"1.2345"`) //string
	floatBlankJSON  = []byte(`"(:"`)     //wrong nullFloatJSON
	nullFloatJSON   = []byte(`{"Float64":1.2345,"Valid":true}`)
)

func TestMarshalZeroValueFloat(t *testing.T) {
	v1, _ := json.Marshal(0)

	x := NullFloat64{sql.NullFloat64{Float64: 0, Valid: true}}
	v2, _ := x.MarshalJSON()

	// Equal to 0 if two byte slices match.
	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNulValueFloat(t *testing.T) {
	v1, _ := json.Marshal(nil)

	x := NullFloat64{sql.NullFloat64{Float64: 0, Valid: false}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalPositiveValueFloat(t *testing.T) {
	v1, _ := json.Marshal(200000)

	x := NullFloat64{sql.NullFloat64{Float64: 200000, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestMarshalNegativeValueFloat(t *testing.T) {
	v1, _ := json.Marshal(-31415926.5)

	x := NullFloat64{sql.NullFloat64{Float64: -31415926.5, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestUnmarshalZeroValueFloat(t *testing.T) {
	v1 := NullFloat64{sql.NullFloat64{Float64: 0, Valid: true}}
	v2 := NullFloat64{}

	x := []byte(strconv.Itoa(0))

	err := v2.UnmarshalJSON(x)
	if err != nil {
		t.Logf("%v", err)
	}

	if v1.Float64 != v2.Float64 || v1.Valid != v2.Valid {
		t.Logf("%v and %v does not match!", v1, v2)
		t.Fail()
	}
}

func TestUnmarshalFloat(t *testing.T) {
	var f NullFloat64
	err := f.UnmarshalJSON(floatJSON)
	if err != nil {
		t.Fail()
	}
	t.Logf("%v", err)

	var sf NullFloat64
	err = sf.UnmarshalJSON(floatStringJSON)
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)

	var nf NullFloat64
	err = nf.UnmarshalJSON(nullFloatJSON)
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)

	var blank NullFloat64
	err = blank.UnmarshalJSON(floatBlankJSON)
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)

	var badType NullFloat64
	err = badType.UnmarshalJSON(boolJSON)
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)

	var invalid NullFloat64
	err = invalid.UnmarshalJSON(invalidJSON)
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)
}

func TestTextUnmarshalFloat(t *testing.T) {
	var f NullFloat64
	err := f.UnmarshalJSON([]byte(`11.2345`))
	if err != nil {
		t.Fail()
	}
	t.Logf("%v", err)

	var blank NullFloat64
	err = blank.UnmarshalJSON([]byte(""))
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)
	var null NullFloat64
	err = null.UnmarshalJSON([]byte("null"))
	if err != nil {
		t.Fail()
	}
	t.Logf("%v", err)
	var invalid NullFloat64
	err = invalid.UnmarshalJSON([]byte("hello world"))
	if err == nil {
		t.Fail()
	}
	t.Logf("%v", err)
}

// func TestMarshalFloat(t *testing.T) {
// 	f := FloatFrom(1.2345)
// 	data, err := json.Marshal(f)

// 	// invalid values should be encoded as null
// 	null := NewFloat(0, false)
// 	data, err = json.Marshal(null)
// }
