package testing

import (
	"fmt"
	"testing"
)

func assertEqual(t testing.TB, a any, b any) {
	if a != b {

		t.Errorf("Well, sorry buddy but: %+v != %+v", a, b)
	}
}

func TestTB(t *testing.T) {
	assertEqual(t, 1, 2)
}

func BenchmarkTB(t *testing.B) {
	assertEqual(t, 1, 1)

	for i := 0; i < t.N; i++ {
		fmt.Println("benchmarking...")
	}
}
