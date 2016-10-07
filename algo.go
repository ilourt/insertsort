package binarysearch

// Search the int
func Search(list []int, value int) bool {
	ln := len(list)

	if ln == 1 {
		if list[0] == value {
			return true
		}

		return false
	}

	var middle int
	if ln%2 == 1 {
		middle = (ln - 1) / 2
	} else {
		middle = ln / 2
	}

	if value < list[middle] {
		return Search(list[:middle], value)
	}

	return Search(list[middle:], value)
}
