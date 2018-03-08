package reverse_string

import "testing"

func TestReverse(t *testing.T) {
	reverse := reverse_string("hello")
	if reverse != "olleh" {
		t.Errorf("Reverse was incorrect, got %s, want: %s", reverse, "olleh")
	} else {
		t.Log("Reverse tested successfully")
	}
}
