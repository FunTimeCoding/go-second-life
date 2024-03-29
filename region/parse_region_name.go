package second_life

import "strings"

func ParseRegionName(region string) string {
	var name string
	position := strings.LastIndex(region, "(")

	if position != -1 {
		name = region[:position-1]
	}

	return name
}
