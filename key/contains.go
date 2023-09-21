package key

func Contains(
	elements []Key,
	search Key,
) bool {
	for _, v := range elements {
		if search == v {
			return true
		}
	}

	return false
}
