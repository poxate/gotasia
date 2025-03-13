package gotasia

import "testing"

type testRange = struct {
	Name       string "json:\"name\""
	RangeEnd   int    "json:\"rangeEnd\""
	RangeStart int    "json:\"rangeStart\""
	Value      any    "json:\"value\""
	ValueType  string "json:\"valueType\""
}

func TestFlattenRange(t *testing.T) {
	attrs := testCreateAttributes()
	attrs.Keyframes[0].Value = []testRange{
		{RangeStart: 0, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 8},
		{RangeStart: 3, RangeEnd: 5},
		{RangeStart: 4, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 1},
		{RangeStart: 5, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 8},
		{RangeStart: 0, RangeEnd: 3},
		{RangeStart: 1, RangeEnd: 4},
	}
	spans := flattenRange("abcdefgh", attrs)
	texts := []string{"a", "bc", "d", "e", "fgh"}

	if len(spans) != len(texts) {
		t.Fatalf("expected %d spans, got: %d", len(texts), len(spans))
	}

	for i, text := range texts {
		if text != spans[i].Text {
			t.Fatalf("expected %s in span %d, got: %s", text, i, spans[i].Text)
		}
	}
}

func testCreateAttributes() rawTextAttributes {
	return rawTextAttributes{Keyframes: []struct {
		EndTime int "json:\"endTime\""
		Time    int "json:\"time\""
		Value   []struct {
			Name       string "json:\"name\""
			RangeEnd   int    "json:\"rangeEnd\""
			RangeStart int    "json:\"rangeStart\""
			Value      any    "json:\"value\""
			ValueType  string "json:\"valueType\""
		} "json:\"value\""
		Duration int "json:\"duration\""
	}{
		{
			Value: []struct {
				Name       string "json:\"name\""
				RangeEnd   int    "json:\"rangeEnd\""
				RangeStart int    "json:\"rangeStart\""
				Value      any    "json:\"value\""
				ValueType  string "json:\"valueType\""
			}{},
		},
	}}
}
