package option

import "testing"

func TestNormalizeFastupPasswordDerivesExactSuffix(t *testing.T) {
	got, enabled := NormalizeFastupPassword("synthetic-password#fastup", "synthetic-mpw")
	if !enabled {
		t.Fatal("Fastup suffix was not detected")
	}
	if got == "synthetic-password#fastup" || len(got) != 32 {
		t.Fatalf("unexpected derived password: %q", got)
	}
	if got != "2879f15592f10270f2aa568ca456779b" {
		t.Fatalf("unexpected Fastup derivation: %q", got)
	}
}

func TestNormalizeFastupPasswordLeavesNearSuffixUntouched(t *testing.T) {
	got, enabled := NormalizeFastupPassword("ordinary#fastup-not-a-suffix", "ignored")
	if enabled || got != "ordinary#fastup-not-a-suffix" {
		t.Fatalf("near suffix must remain unchanged: %q, %v", got, enabled)
	}
}
