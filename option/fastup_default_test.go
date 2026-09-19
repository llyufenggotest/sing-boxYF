package option

import "testing"

func TestNormalizeFastupPasswordUsesDefaultMpw(t *testing.T) {
	got, enabled := NormalizeFastupPassword("synthetic-password#fastup", "")
	if !enabled || len(got) != 32 {
		t.Fatalf("default mpw must derive a Fastup password: %q, %v", got, enabled)
	}
}
