package short_url

import "testing"

func TestCode(t *testing.T) {
	id := int64(1100000)
	code := GenCode(id)
	t.Logf("gen code: %v", code)
	parseId, err := ParseCode(code)
	if err != nil {
		t.Error(err)
	}
	if parseId != id {
		t.Fatalf("want %d, got %d", id, parseId)
	}
	t.Logf("parse code: %v", id)
}
