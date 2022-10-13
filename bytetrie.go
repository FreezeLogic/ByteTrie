package bytetrie

type node struct {
	items []*item
}
type item struct {
	val   byte
	last  bool
	child *node
}

func NewTrie() *node {
	n := new(node)
	n.items = *new([]*item)
	return n
}

func (n *node) Insert(b []byte) {
	node := n
	tmpitem := new(item)
	for _, vb := range b {
		newitem := node.isExist(vb)
		if newitem == nil {
			node.items = append(node.items, &item{
				val:   vb,
				last:  false,
				child: NewTrie(),
			})
			tmpitem = node.items[len(node.items)-1]
			node = tmpitem.child
		} else {
			tmpitem = newitem
			node = tmpitem.child
		}
	}
	tmpitem.last = true
}

func (n *node) Search(b []byte) bool {
	node := n
	tmpitem := new(item)
	for _, vb := range b {
		item := node.isExist(vb)
		if item == nil {
			return false
		} else {
			tmpitem = item
			node = tmpitem.child
		}
	}
	if tmpitem.last {
		return true
	}
	return false
}

func (n *node) isExist(b byte) *item {
	for _, v := range n.items {
		if v.val == b {
			return v
		}
	}
	return nil
}
