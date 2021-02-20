package numutils

import "version"

var v = version.Version

func Add(a, b int) int {
	return a + b
}

func GetVersion() string {
	return v
}
