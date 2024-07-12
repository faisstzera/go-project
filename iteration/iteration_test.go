package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaaa"

	repeated_len := len(repeated)

	expected_len := len(expected)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

	if repeated_len != expected_len {
		t.Errorf("expected %d but got %d", expected_len, repeated_len)
	}
}

func ExampleRepeat() {
	repeated_string := Repeat("a")
	fmt.Println(repeated_string)
	// Output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
