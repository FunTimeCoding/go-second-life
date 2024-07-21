package alias

type Button string

func (b Button) String() string {
	return string(b)
}
