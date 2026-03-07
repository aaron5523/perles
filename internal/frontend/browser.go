package frontend

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// OpenBrowser attempts to open the given URL in the default browser.
// Returns an error if it fails, allowing the caller to fall back to
// printing the URL to the terminal.
func OpenBrowser(url string) error {
	// Check for BROWSER env var first (common on Linux)
	if browser := os.Getenv("BROWSER"); browser != "" {
		return exec.CommandContext(context.Background(), browser, url).Start() //nolint:gosec // BROWSER is user-controlled env var
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.CommandContext(context.Background(), "open", url) //nolint:gosec // url is an internally generated localhost URL
	case "linux":
		cmd = exec.CommandContext(context.Background(), "xdg-open", url) //nolint:gosec // url is an internally generated localhost URL
	case "windows":
		cmd = exec.CommandContext(context.Background(), "cmd", "/c", "start", url) //nolint:gosec // url is an internally generated localhost URL
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Start()
}
