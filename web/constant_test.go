package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Integer(t, 4096, MediumResponseThreshold)
	assert.Integer(t, 2048, MaximumRequestSize)

	assert.Strings(
		t,
		[]string{
			"<h1>Service Unavailable</h1><p>Your request is temporarily unable to be fulfilled.</p>",
			"<h1>Internal Server Error</h1><p>Your request failed to rez.</p>",
		},
		[]string{ServiceUnavailable, InternalServerError},
	)

	assert.Strings(
		t,
		[]string{
			"X-SecondLife-Object-Key",
			"X-SecondLife-Object-Name",
			"X-SecondLife-Owner-Key",
			"X-SecondLife-Owner-Name",
			"X-SecondLife-Region",
			"X-SecondLife-Local-Position",
		},
		[]string{
			ObjectKeyHeader,
			ObjectNameHeader,
			OwnerKeyHeader,
			OwnerNameHeader,
			RegionHeader,
			PositionHeader,
		},
	)
}
