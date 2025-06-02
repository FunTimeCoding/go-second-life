package agent

import (
	"regexp"
	"strings"
)

func FirstName(
	display string,
	legacy string,
) string {
	if regexp.MustCompile(`^[a-zA-Z\d ]+$`).MatchString(display) {
		space := strings.Index(display, " ")

		if space != -1 {
			return display[:space]
		}

		return display
	}

	if strings.HasSuffix(legacy, ResidentSuffix) {
		return strings.TrimSuffix(legacy, ResidentSuffix)
	}

	space := strings.Index(legacy, " ")

	if space != -1 {
		return legacy[:space]
	}

	return legacy
}
