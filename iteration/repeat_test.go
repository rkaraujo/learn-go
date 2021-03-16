package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertResult := func(t testing.TB, repeated string, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertResult(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
