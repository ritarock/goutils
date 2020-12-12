package random

import "testing"

func TestGetUuid(t *testing.T) {
	uuid1 := GetUuid()
	uuid2 := GetUuid()
	result := uuid1 == uuid2
	expect := false
	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}
