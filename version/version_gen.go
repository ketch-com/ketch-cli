// Code generated by winch. DO NOT EDIT.
package version

import (
	"fmt"
	"runtime"
)

const (
	Name        = "ketch"
	Description = "Ketch CLI is the command-line interface to ketch"
	ReleaseName = "testy coyote"
	Version     = "0.0.0"
	Prerelease  = "dev"
)

// String returns the complete version string, including prerelease
func String() string {
	s := fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s %s", Version, Prerelease, s)
	}
	return fmt.Sprintf("%s %s", Version, s)
}
