package tst

type TSTIterator struct {
	tst   *TernarySearchTrie
	stack []iteratorFrame
	ready bool
}

type iteratorFrame struct {
	node   *Node
	prefix []rune
}

func NewTSTIterator(tst *TernarySearchTrie) *TSTIterator {
	iter := &TSTIterator{tst: tst, stack: make([]iteratorFrame, 0), ready: false}
	iter.stack = append(iter.stack, iteratorFrame{node: iter.tst.root, prefix: make([]rune, 0)})
	iter.advance()
	return iter
}

func (iter *TSTIterator) advance() {
	for len(iter.stack) > 0 {
		frame := iter.stack[len(iter.stack)-1]
		node := frame.node
		iter.stack = iter.stack[:len(iter.stack)-1]

		if node.right != nil {
			iter.stack = append(iter.stack, iteratorFrame{node: node.right, prefix: frame.prefix})
		}

		if node.middle != nil {
			newPrefix := append(frame.prefix, node.char)
			iter.stack = append(iter.stack, iteratorFrame{node: node.middle, prefix: newPrefix})
		}

		if node.left != nil {
			iter.stack = append(iter.stack, iteratorFrame{node: node.left, prefix: frame.prefix})
		}

		if node.value != 0 {
			iter.stack = append(iter.stack, frame)
			iter.ready = true
			return
		}
	}
	iter.ready = false
}

func (iter *TSTIterator) HasNext() bool {
	return iter.ready
}

func (iter *TSTIterator) Next() (string, interface{}, bool) {
	if !iter.ready {
		return "", nil, false
	}

	frame := iter.stack[len(iter.stack)-1]
	key := string(append(frame.prefix, frame.node.char))
	value := frame.node.value
	iter.stack = iter.stack[:len(iter.stack)-1]
	iter.advance()

	return key, value, true
}
