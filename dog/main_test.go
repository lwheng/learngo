package dog

import (
	"testing"
)

func TestYears(t *testing.T) {
	if x := Years(2); x != 14 {
		t.Error("Expected", 14, "got", x)
	}
}
