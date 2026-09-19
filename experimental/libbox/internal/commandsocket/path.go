package commandsocket

import (
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
)

// Path returns a Unix socket path below the shared container. Darwin limits
// Unix-domain socket pathnames to 104 bytes.
func Path(basePath string) string {
	path := filepath.Join(basePath, "command.sock")
	if len([]byte(path)) < 104 {
		return path
	}
	hash := sha256.Sum256([]byte(basePath))
	return filepath.Join(filepath.Dir(basePath), "c"+hex.EncodeToString(hash[:5])+".sock")
}
