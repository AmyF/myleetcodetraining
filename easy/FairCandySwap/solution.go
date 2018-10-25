func fairCandySwap(A []int, B []int) []int {
	suma, sumb := 0, 0
	for _, v := range A {
		suma += v
	}
	for _, v := range B {
		sumb += v
	}
	delta := (sumb - suma) / 2
	m := map[int]int{}
	for _, v := range B {
		m[v] = v
	}
	for _, v := range A {
		if m[v+delta] > 0 {
			return []int{v, v + delta}
		}
	}
	return []int{}
}