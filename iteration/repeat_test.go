package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestCounter(t *testing.T) {
	counted := Count("a", "aaaaa")
	expected := 5

	if counted != expected {
		t.Errorf("expected %d but got %d", expected, counted)
	}
}

func TestIndex(t *testing.T) {
	index := Index("a", "qwertayuiop")
	expected := 5

	if index != expected {
		t.Errorf("expected %d but got %d", expected, index)
	}
}

func TestJoin(t *testing.T) {
	days := []string{"Monday", "Tuesday", "Wednesday"}
	joined := Join(days, ", ")
	expected := "Monday, Tuesday, Wednesday"

	if joined != expected {
		t.Errorf("expected %q but got %q", expected, joined)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}
