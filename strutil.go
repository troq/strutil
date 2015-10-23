package strutil

func StringInSlice(str string, slice []string) bool {
	for _, el := range slice {
		if str == el {
			return true
		}
	}
	return false
}

func StringIndex(str string, slice []string) int {
	for i, s := range slice {
		if s == str {
			return i
		}
	}
	return -1
}

func Uniq(slice []string) (uslice []string) {
	for _, s := range slice {
		if !StringInSlice(s, uslice) {
			uslice = append(uslice, s)
		}
	}
	return
}
