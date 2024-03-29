package region

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestParseRegionName(t *testing.T) {
	assertParseRegionName(
		t,
		"Example (224792, 333834)",
		"Example",
	)
}

func assertParseRegionName(
	t *testing.T,
	region string,
	expected string,
) {
	t.Helper()
	assert.String(t, expected, ParseRegionName(region))
}
