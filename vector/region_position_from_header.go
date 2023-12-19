package vector

func RegionPositionFromHeader(position string) *Vector {
	return FromStringCustom(position, "()").Round(1)
}
