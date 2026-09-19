package libbox

import (
 "path/filepath"
 "strings"
 "testing"
)

func TestCommandSocketPathKeepsShortBase(t *testing.T) {
 base := "/tmp/singbox"
 if got := commandSocketPath(base); got != filepath.Join(base, "command.sock") {
  t.Fatalf("short path changed: %q", got)
 }
}

func TestCommandSocketPathShortensLongAppGroup(t *testing.T) {
 base := "/private/var/mobile/Containers/Shared/AppGroup/8ABC3121-2825-4F54-91DC-98A307ECE71A/group.io.llyufenggotest.singboxyf"
 got := commandSocketPath(base)
 if len([]byte(got)) >= 104 {
  t.Fatalf("socket path too long for Darwin: %d bytes: %q", len([]byte(got)), got)
 }
 if !strings.HasPrefix(got, filepath.Dir(base)+"/") {
  t.Fatalf("socket escaped shared container: %q", got)
 }
 if got != commandSocketPath(base) {
  t.Fatal("server and client paths must match")
 }
 if got == commandSocketPath(base+"-another") {
  t.Fatal("different app groups must not share a socket")
 }
}
