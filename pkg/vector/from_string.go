package vector

func FromString(v string) *Vector {
	return FromStringCustom(v, "<>")
}
