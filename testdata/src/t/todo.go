package t

func inc(i *int) {
	*i++
}

func whatIfPassToFunction(kCount *int) {
	inc(kCount) // can we check the parameter name to know const is assigned to non-const?
}

func whatIfAssign(kMap map[int]int) {
	m := kMap // also warn here?
	_ = m
}
