package t

func failMap(kInts, ints map[int]int) {
	ints[0] = 0
	kInts[0] = 0             // want "write to const variable 'kInts'"
	kInts[1], ints[1] = 1, 1 // want "write to const variable 'kInts'"
}

func failSlice(bs, kBs []byte) {
	kBs[0] = 0 // want "write to const variable 'kBs'"
	bs[0] = 0
}

type myStruct struct {
	i int
}

func failStruct(kStruct, s *myStruct) {
	s.i, kStruct.i = 1, 1 // want "write to const variable 'kStruct'"
	(*kStruct).i = 2      // want "write to const variable 'kStruct'"
}

type myNode struct {
	n     *myNode
	kData int
}

func failMultiLevelNested(n, kNode *myNode) {
	n.kData = 1   // want "write to const variable 'n.kData'"
	n.n.kData = 2 // want "write to const variable 'n.n.kData'"
	kNode.n = nil // want "write to const variable 'kNode'"
	kNode.n.n = nil
	kNode.n.n.kData = 3 // want "write to const variable 'kNode.n.n.kData'"

	*kNode = myNode{} // want "write to const variable 'kNode'"
}

func failInc(kCount, count *int) {
	*kCount++    // want "write to const variable 'kCount'"
	*kCount -= 1 // want "write to const variable 'kCount'"
	*count--
	*count *= 2
}
