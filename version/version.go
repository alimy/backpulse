package version

import "github.com/coreos/go-semver/semver"

var (
	// Version is the specification version that the package types support.
	Version = semver.Version{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}

	// GitHash Value will be set during build
	GitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)
