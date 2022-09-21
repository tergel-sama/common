package types

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strconv"
	"testing"
)

var (
	boolJSON     = []byte(`true`)
	falseJSON    = []byte(`false`)
	nullBoolJSON = []byte(`{"Bool":true,"Valid":true}`)
	invalidJSON  = []byte(`{"Bool":true,"Valid":false}`)
)

// энэ тзст дараах хэлбэрээр дата оруулна
// 1. Bool null утга шалгах
// 2. Null утгыг оруулж шалгах
// 3. Зөв утгаар шалгах
// Тус бүрийг нь 0 утгатай тэнцэж байгаа эсэхийг шалгана
// Marshal ба Unmarshal дээр тус бүр зөв ажиллаж байгаа эсэхийг шалгана

func TestBoolZeroValue(t *testing.T) {
	v1, _ := json.Marshal(false)

	x := NullBool{sql.NullBool{Bool: false, Valid: true}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

func TestBoolNulValue(t *testing.T) {
	v1, _ := json.Marshal(nil)

	x := NullBool{sql.NullBool{Bool: false, Valid: false}}
	v2, _ := x.MarshalJSON()

	if bytes.Compare(v1, v2) != 0 {
		t.Fail()
	}
}

// func TestMarshalTValue(t *testing.T) {
// 	v1, _ := json.Marshal(true)

// 	x := NullBool{sql.NullBool{Bool: true, Valid: true}}
// 	v2, _ := x.MarshalJSON()

// 	if bytes.Compare(v1, v2) != 0 {
// 		t.Fail()
// 	}
// }

func TestUnmarshal(t *testing.T) {
	v1 := NullBool{sql.NullBool{Bool: false, Valid: false}}
	v2 := NullBool{}

	x := []byte(strconv.Itoa(0))

	err := v2.UnMarshalJSON(x)
	if err == nil {
		t.Logf("%v", err)
		t.Fail()
	}

	if v1.Bool != v2.Bool || v1.Valid != v2.Valid {
		t.Logf("%v and %v does not match!", v1, v2)
		t.Fail()
	}
}

func TestUnmarshalBool(t *testing.T) {
	var b NullBool
	err := json.Unmarshal(boolJSON, &b.NullBool)
	if err != nil {
		t.Logf("%v", err)
	}

	var nb NullBool
	err = json.Unmarshal(nullBoolJSON, &nb)
	if err != nil {
		panic("err should not be nil")
	}

	var null NullBool
	err = json.Unmarshal(nullBoolJSON, &null)
}

func TestMarshalBool(t *testing.T) {
	b := NullBool{sql.NullBool{Bool: true, Valid: true}}
	_, err := b.MarshalJSON()
	if err != nil {
		t.Fail()
	}
}
