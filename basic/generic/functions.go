package generic

func ToMap[T comparable](slice []T) map[T]struct{} {
	ret := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		ret[v] = struct{}{}
	}
	return ret
}

func RemoveDup[T comparable](slice []T) []T {
	m := ToMap(slice)
	ret := make([]T, 0, len(m))
	for v := range m {
		ret = append(ret, v)
	}
	return ret
}
