package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//go:embed index.html
var assets embed.FS

func main() {
	htmlBytes, err := assets.ReadFile("index.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load bundled index.html: %v\n", err)
		os.Exit(1)
	}

	tempDir, err := os.MkdirTemp("", "blem-")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create temp directory: %v\n", err)
		os.Exit(1)
	}

	htmlPath := filepath.Join(tempDir, "index.html")
	if err := os.WriteFile(htmlPath, htmlBytes, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write temporary index.html: %v\n", err)
		os.Exit(1)
	}

	if err := openInBrowser(htmlPath); err != nil {
		fmt.Fprintf(os.Stderr, "failed to open Blem in your browser: %v\n", err)
		os.Exit(1)
	}
}

func openInBrowser(target string) error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", target)
		return cmd.Start()
	case "darwin":
		return exec.Command("open", target).Start()
	default:
		return exec.Command("xdg-open", target).Start()
	}
}
