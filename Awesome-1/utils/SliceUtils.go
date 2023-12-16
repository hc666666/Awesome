package utils

import "github.com/google/gopacket"

func InMap[T gopacket.LayerType](m map[T]struct{}, t T) bool {
	_, ok := m[t]
	return ok
}
func ConvertSlice2Map[T gopacket.LayerType](sl []T) map[T]struct{} {
	set := make(map[T]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}
