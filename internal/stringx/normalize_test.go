package stringx

import "testing"

func TestNormalizeEmail(t *testing.T) {
	got := NormalizeEmail(" USER@EXAMPLE.COM ")
	want := "user@example.com"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
