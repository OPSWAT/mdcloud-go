package utils

// StringInSlice checks for string in slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// StringSlice makes a slice str to str pointers
func StringSlice(strs []string) []*string {
	res := make([]*string, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = &(strs[i])
	}
	return res
}
