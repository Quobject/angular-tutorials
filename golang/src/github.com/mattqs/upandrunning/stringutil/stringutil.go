package stringutil

//FullNameNakedReturn a
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}

//FullName lll
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}
