package web

const (
	MaximumResponseSize     = int(16384)
	LargeResponseThreshold  = MaximumResponseSize / 2
	MediumResponseThreshold = LargeResponseThreshold / 2
	MaximumRequestSize      = int(2048)
)

const (
	ServiceUnavailable  = "<h1>Service Unavailable</h1><p>Your request is temporarily unable to be fulfilled.</p>"
	InternalServerError = "<h1>Internal Server Error</h1><p>Your request failed to rez.</p>"
)

const (
	ObjectKeyHeader  = "X-SecondLife-Object-Key"
	ObjectNameHeader = "X-SecondLife-Object-Name"
	OwnerKeyHeader   = "X-SecondLife-Owner-Key"
	OwnerNameHeader  = "X-SecondLife-Owner-Name"
	RegionHeader     = "X-SecondLife-Region"
	PositionHeader   = "X-SecondLife-Local-Position"
)
