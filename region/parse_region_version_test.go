package second_life

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestParseRegionVersion(t *testing.T) {
	assertParseRegionVersion(
		t,
		"Second-Life-LSL/2022-03-24.569934 (https://secondlife.com)",
		"2022-03-24.569934",
	)
}

func assertParseRegionVersion(
	t *testing.T,
	userAgent string,
	expected string,
) {
	t.Helper()
	assert.String(t, ParseRegionVersion(userAgent), expected)
}
