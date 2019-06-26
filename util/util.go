package util

// Sunday sunday
func Sunday(text, sub string) bool {
	if len(sub) == 0 {
		return true
	}

	flag := false
	subLast := len(sub) - 1
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == sub[subLast] {
			flag = true
			subLast--
			// sub remaining length is less than 0, no need to compare
			if subLast < 0 {
				break
			}
		} else {
			// no match has been successful yet
			if flag == false {
				if i+1 < len(sub) {
					break
				}
			} else {
				// once matched successfully but can't match continuously
				flag = false
				break
			}
		}
	}

	return flag
}
