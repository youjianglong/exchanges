package common

import (
	"encoding/json"
	"testing"
)

func TestStrictUnmarshal_EEConflict(t *testing.T) {
	type event struct {
		Event string `json:"e"`
		Time  Int64  `json:"E"`
	}

	var e event
	err := StrictUnmarshal([]byte(`{"e":"miniTicker","E":1700000000000}`), &e)
	if err != nil {
		t.Fatal(err)
	}
	if e.Event != "miniTicker" {
		t.Fatalf("Event = %q, want miniTicker", e.Event)
	}
	if e.Time.Value() != 1700000000000 {
		t.Fatalf("Time = %d, want 1700000000000", e.Time.Value())
	}
}

func TestStrictUnmarshal_Int64FromNumber(t *testing.T) {
	type s struct {
		Time Int64 `json:"E"`
	}
	var v s
	if err := StrictUnmarshal([]byte(`{"E":123}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.Time.Value() != 123 {
		t.Fatalf("Time = %d, want 123", v.Time.Value())
	}
}

func TestStrictUnmarshal_Float64FromString(t *testing.T) {
	type s struct {
		Price Float64 `json:"c"`
	}
	var v s
	if err := StrictUnmarshal([]byte(`{"c":"123.45"}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.Price.String() != "123.45" {
		t.Fatalf("Price = %q, want 123.45", v.Price.String())
	}
}

func TestStrictUnmarshal_RawMessage(t *testing.T) {
	type envelope struct {
		Data json.RawMessage `json:"data"`
	}
	raw := []byte(`{"data":{"e":"ticker","E":1}}`)
	var e envelope
	if err := StrictUnmarshal(raw, &e); err != nil {
		t.Fatal(err)
	}
	want := `{"e":"ticker","E":1}`
	if string(e.Data) != want {
		t.Fatalf("Data = %s, want %s", e.Data, want)
	}
}

func TestStrictDecode_Slice(t *testing.T) {
	type item struct {
		Event string `json:"e"`
		Time  Int64  `json:"E"`
	}
	var items []*item
	err := StrictDecode([]byte(`[{"e":"a","E":1},{"e":"b","E":2}]`), &items)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("len = %d, want 2", len(items))
	}
	if items[0].Event != "a" || items[1].Event != "b" {
		t.Fatalf("events = %q, %q", items[0].Event, items[1].Event)
	}
}

func TestStrictDecode_CamelCase(t *testing.T) {
	type resp struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	}
	var r resp
	if err := StrictDecode([]byte(`{"code":"0","msg":"ok"}`), &r); err != nil {
		t.Fatal(err)
	}
	if r.Code != "0" || r.Msg != "ok" {
		t.Fatalf("resp = %+v", r)
	}
}

func TestStrictDecode_EmptySlice(t *testing.T) {
	var items []struct {
		Symbol string `json:"symbol"`
	}
	if err := StrictDecode([]byte(`[]`), &items); err != nil {
		t.Fatal(err)
	}
	if len(items) != 0 {
		t.Fatalf("len = %d, want 0", len(items))
	}
}

func TestStrictUnmarshal_StringTagOption(t *testing.T) {
	type apiErr struct {
		Code int64 `json:"code,string"`
	}
	var e apiErr
	if err := StrictUnmarshal([]byte(`{"code":"50103"}`), &e); err != nil {
		t.Fatal(err)
	}
	if e.Code != 50103 {
		t.Fatalf("Code = %d, want 50103", e.Code)
	}
}

func TestStrictDecode_NestedArrayFallback(t *testing.T) {
	var rows [][]Mixed
	if err := StrictDecode([]byte(`[["1","2"],["3","4"]]`), &rows); err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[0][0].String() != "1" {
		t.Fatalf("rows = %+v", rows)
	}
}
