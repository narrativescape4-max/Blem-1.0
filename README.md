# Blem-1.0

A 3D software to play around with.

## Build a Windows `.exe`

This repo includes a small Go launcher that embeds `index.html` and opens it in your default browser.

> Note: Binary files are not committed to this repository. Build artifacts are generated locally in `dist/`.

Use the build script:

```bash
./scripts/build_windows_exe.sh
```

Or run the command directly:

```bash
GOOS=windows GOARCH=amd64 go build -o dist/Blem.exe .
```

The built file will be created at `dist/Blem.exe`.
