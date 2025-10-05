package main

import "github.com/funtimecoding/go-second-life/pkg/sentry"

func main() {
	if true {
		return
	}

	sentry.SetObjectDetails(nil, "", "", "")
	sentry.SetRegionDetails(nil, "", "", nil)
}
