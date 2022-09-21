package types

import (
	"testing"
)

var (
	timeString1  = "2012-12-21T21:21:21Z"
	timeString2  = "2012-12-21T22:21:21+01:00" // Same time as timeString1 but in a different timezone
	timeString3  = "2018-08-19T01:02:03Z"
	timeJSON     = []byte(`"` + timeString1 + `"`)
	nullTimeJSON = []byte(`null`)
	timeObject   = []byte(`{"Time":"2012-12-21T21:21:21Z","Valid":true}`)
	nullObject   = []byte(`{"Time":"0001-01-01T00:00:00Z","Valid":false}`)
	badObject    = []byte(`{"hello": "world"}`)
)

func TestUnmarshalJSONTimeJSON(t *testing.T) {
	var ti NullTime
	err := ti.UnmarshalJSON(timeJSON)
	if err != nil {
		t.Fail()
	}

	var null NullTime
	err = null.UnmarshalJSON(nullTimeJSON)

	if err != nil {
		t.Fail()
	}
	var fromObject NullTime
	err = fromObject.UnmarshalJSON(timeObject)

	if err == nil {
		t.Fail()
	}
	var nullFromObj NullTime
	err = nullFromObj.UnmarshalJSON(nullObject)
	if err == nil {
		t.Fail()
	}
	var bad NullTime
	err = bad.UnmarshalJSON(badObject)
	if err == nil {
		t.Fail()
	}
}
