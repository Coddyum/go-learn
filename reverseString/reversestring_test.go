package reversestring

import "testing"


func TestReverseString(t *testing.T){
	got := Name("Coddy")
	want := "yddoC"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func BenchmarkName(b *testing.B) {
	for b.Loop() {
		Name("Coddy")
	}
}