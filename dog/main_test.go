package dog

import (
	"testing"
)

func TestYears(t *testing.T) {
	if x := Years(2); x != 14 {
		t.Error("Expected", 14, "got", x)
	}
}

func TestYearsTwo(t *testing.T) {
	if x := YearsTwo(2); x != 14 {
		t.Error("Expected", 14, "got", x)
	}
}

func BenchmarkYears(b *testing.B) {
	Years(2)
}

func BenchmarkYearsTwo(b *testing.B) {
	YearsTwo(2)
}

func ExampleYears() {
	Years(2)
	// Output:
	// 14
}

func ExampleYearsTwo() {
	YearsTwo(2)
	// Output:
	// 14
}
