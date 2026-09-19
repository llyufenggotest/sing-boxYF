package libbox

import (
 "crypto/sha256"
 "encoding/hex"
 "path/filepath"
)

// Darwin's Unix-domain socket pathname limit is 104 bytes. The app-group
// compatibility shim may add a long directory below the shared container.
func commandSocketPath(basePath string) string {
 path := filepath.Join(basePath, "command.sock")
 if len([]byte(path)) < 104 {
  return path
 }
 hash := sha256.Sum256([]byte(basePath))
 return filepath.Join(filepath.Dir(basePath), "c"+hex.EncodeToString(hash[:5])+".sock")
}
