package commandsocket

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPathKeepsShortBase(t *testing.T) {
	base := "/tmp/singbox"
	if got := Path(base); got != filepath.Join(base, "command.sock") {
		t.Fatalf("short path changed: %q", got)
	}
}

func TestPathShortensLongAppGroup(t *testing.T) {
	base := "/private/var/mobile/Containers/Shared/AppGroup/8ABC3121-2825-4F54-91DC-98A307ECE71A/group.io.llyufenggotest.singboxyf"
	got := Path(base)
	if len([]byte(got)) >= 104 {
		t.Fatalf("socket path too long for Darwin: %d bytes: %q", len([]byte(got)), got)
	}
	if !strings.HasPrefix(got, filepath.Dir(base)+"/") {
		t.Fatalf("socket escaped shared container: %q", got)
	}
	if got != Path(base) {
		t.Fatal("server and client paths must match")
	}
	if got == Path(base+"-another") {
		t.Fatal("different app groups must not share a socket")
	}
}
