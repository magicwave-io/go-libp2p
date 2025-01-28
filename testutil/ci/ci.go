// Package ci implements some helper functions to use during
// tests. Many times certain facilities are not available, or tests
// must run differently.
package ci

import (
	"os"

	jenkins "github.com/ipfs/go-libp2p/testutil/ci/jenkins"
	travis "github.com/ipfs/go-libp2p/testutil/ci/travis"
)

// EnvVar is a type to use travis-only env var names with
// the type system.
type EnvVar string

// Environment variables that TravisCI uses.
const (
	VarCI      EnvVar = "CI"
	VarNoFuse  EnvVar = "TEST_NO_FUSE"
	VarVerbose EnvVar = "TEST_VERBOSE"
)

// IsRunning attempts to determine whether this process is
// running on CI. This is done by checking any of:
//
//  CI=true
//  travis.IsRunning()
//  jenkins.IsRunning()
//
func IsRunning() bool {
	if os.Getenv(string(VarCI)) == "true" {
		return true
	}

	return travis.IsRunning() || jenkins.IsRunning()
}

// Env returns the value of a CI env variable.
func Env(v EnvVar) string {
	return os.Getenv(string(v))
}

// Returns whether FUSE is explicitly disabled wiht TEST_NO_FUSE.
func NoFuse() bool {
	return os.Getenv(string(VarNoFuse)) == "1"
}

// Returns whether TEST_VERBOSE is enabled.
func Verbose() bool {
	return os.Getenv(string(VarVerbose)) == "1"
}
