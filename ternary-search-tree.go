package tst

type Node struct {
	char   rune
	left   *Node
	middle *Node
	right  *Node
	value  int
}

func NewNode(char rune) *Node {
	return &Node{
		char:   char,
		left:   nil,
		middle: nil,
		right:  nil,
		value:  0,
	}
}

type TernarySearchTrie struct {
	root *Node
}

func NewTST() *TernarySearchTrie {
	return &TernarySearchTrie{
		root: nil,
	}
}

func (tst *TernarySearchTrie) Put(key string, value int) {
	tst.root = tst.put(tst.root, strToRuneSlice(key), value, 0)
}

func strToRuneSlice(key string) []rune {
	return []rune(key)
}

func (tst *TernarySearchTrie) put(node *Node, key []rune, value int, index int) *Node {
	c := key[index]

	if node == nil {
		node = NewNode(c)
	}

	if c < node.char {
		node.left = tst.put(node.left, key, value, index)
	} else if c > node.char {
		node.right = tst.put(node.right, key, value, index)
	} else if index < len(key)-1 {
		node.middle = tst.put(node.middle, key, value, index+1)
	} else {
		node.value = value
	}

	return node
}

func (tst *TernarySearchTrie) Get(key string) int {
	node := tst.get(tst.root, strToRuneSlice(key), 0)
	if node == nil {
		return -1
	}
	return node.value
}

func (tst *TernarySearchTrie) get(node *Node, key []rune, index int) *Node {
	if node == nil {
		return nil
	}

	c := key[index]

	if c < node.char {
		return tst.get(node.left, key, index)
	} else if c > node.char {
		return tst.get(node.right, key, index)
	} else if index < len(key)-1 {
		return tst.get(node.middle, key, index+1)
	} else {
		return node
	}
}

// Deletion
func (tst *TernarySearchTrie) Delete(key string) {
	tst.root = tst.delete(tst.root, strToRuneSlice(key), 0)
}

func (tst *TernarySearchTrie) delete(node *Node, key []rune, index int) *Node {
	if node == nil {
		return nil
	}

	c := key[index]

	if c < node.char {
		node.left = tst.delete(node.left, key, index)
	} else if c > node.char {
		node.right = tst.delete(node.right, key, index)
	} else {
		if index == len(key)-1 {
			node.value = 0
		} else {
			node.middle = tst.delete(node.middle, key, index+1)
		}
	}

	if node.left == nil && node.middle == nil && node.right == nil && node.value == 0 {
		return nil
	}

	return node
}

// Prefix search
func (tst *TernarySearchTrie) KeysWithPrefix(prefix string) []string {
	var results []string
	node := tst.get(tst.root, strToRuneSlice(prefix), 0)

	if node == nil {
		return results
	}

	if node.value != 0 {
		results = append(results, prefix)
	}

	return tst.collect(node.middle, prefix, &results)
}

func (tst *TernarySearchTrie) collect(node *Node, prefix string, results *[]string) []string {
	if node == nil {
		return *results
	}

	tst.collect(node.left, prefix, results)

	if node.value != 0 {
		*results = append(*results, prefix+string(node.char))
	}

	tst.collect(node.middle, prefix+string(node.char), results)
	tst.collect(node.right, prefix, results)

	return *results
}

func (tst *TernarySearchTrie) RangeCollect(min, max string) []string {
	var keys []string
	tst.rangeCollect(tst.root, "", min, max, &keys)
	return keys
}

func (tst *TernarySearchTrie) rangeCollect(node *Node, prefix string, min, max string, keys *[]string) {
	if node == nil {
		return
	}

	current := prefix + string(node.char)

	if current > min {
		tst.rangeCollect(node.left, prefix, min, max, keys)
	}

	if current >= min && current <= max {
		if node.value != 0 {
			*keys = append(*keys, current)
		}
		tst.rangeCollect(node.middle, current, min, max, keys)
	}

	if current < max {
		tst.rangeCollect(node.right, prefix, min, max, keys)
	}
}

func (tst *TernarySearchTrie) AllKeys() []string {
	var results []string
	return tst.collect(tst.root, "", &results)
}
