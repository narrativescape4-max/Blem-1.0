#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_DIR="$ROOT_DIR/dist"
OUT_FILE="$OUT_DIR/Blem.exe"

mkdir -p "$OUT_DIR"

echo "Building Windows executable..."
(
  cd "$ROOT_DIR"
  GOOS=windows GOARCH=amd64 go build -o "$OUT_FILE" .
)

echo "Build complete: $OUT_FILE"

# Quick signature check (PE files begin with MZ)
python3 - <<'PY'
from pathlib import Path
p = Path('dist/Blem.exe')
header = p.read_bytes()[:2]
if header != b'MZ':
    raise SystemExit(f"unexpected executable header: {header!r}")
print(f"Verified PE header: {header!r}")
PY
