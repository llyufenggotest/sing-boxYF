#!/usr/bin/env bash
set -euo pipefail
APP="${1:?usage: inject_sideload_dylib.sh <SFI.app> <dylib>}"
DYLIB="${2:?usage: inject_sideload_dylib.sh <SFI.app> <dylib>}"
NAME="$(basename "$DYLIB")"
OUTPUT="${3:-$APP}"
EXT="$OUTPUT/PlugIns/Extension.appex"
if [[ ! -d "$EXT" ]]; then EXT="$OUTPUT/PlugIns/SystemExtension.appex"; fi
[[ -d "$EXT" ]] || { echo "error: packet tunnel extension not found" >&2; exit 1; }
[[ -f "$DYLIB" ]] || { echo "error: missing dylib: $DYLIB" >&2; exit 1; }
mkdir -p "$OUTPUT/Frameworks" "$EXT/Frameworks"
cp -f "$DYLIB" "$OUTPUT/Frameworks/$NAME"
cp -f "$DYLIB" "$EXT/Frameworks/$NAME"
chmod 755 "$OUTPUT/Frameworks/$NAME" "$EXT/Frameworks/$NAME"
printf 'injected %s into host and packet tunnel\n' "$NAME"
