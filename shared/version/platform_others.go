//go:build !linux
// +build !linux

package version

func getPlatformVersionStrings() []string {
	return []string{}
}
