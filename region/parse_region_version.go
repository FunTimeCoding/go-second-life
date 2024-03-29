package region

import "strings"

func ParseRegionVersion(userAgent string) string {
	var version string

	start := strings.Index(userAgent, "/")
	end := strings.Index(userAgent, " ")

	if start != -1 && end != -1 {
		version = userAgent[start+1 : end]
	}

	return version
}
