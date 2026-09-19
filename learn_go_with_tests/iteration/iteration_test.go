package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	sum := Repeat("a", 5)
	fmt.Println(sum)
	// Output: aaaaa
}
func TestContainsThisChar(t *testing.T) {
	want := true
	got := ContainsThisChar("a", "team")
	if want != got {
		t.Errorf("wanted %t got %t", want, got)
	}
}
