package strutil

func StringInSlice(str string, slice []string) bool {
	for _, el := range slice {
		if str == el {
			return true
		}
	}
	return false
}

func Uniq(slice []string) (uslice []string) {
	for _, s := range slice {
		if !StringInSlice(s, uslice) {
			uslice = append(uslice, s)
		}
	}
	return
}
