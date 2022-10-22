package bytetrie

type Node struct {
	items []*Item
}
type Item struct {
	val   byte
	last  bool
	child *Node
}

func New() *Node {
	n := new(Node)
	n.items = *new([]*Item)
	return n
}

func (n *Node) Insert(b []byte) {
	node := n
	tmpitem := new(Item)
	for _, vb := range b {
		newitem := node.isExist(vb)
		if newitem == nil {
			node.items = append(node.items, &Item{
				val:   vb,
				last:  false,
				child: New(),
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

func (n *Node) Search(b []byte) bool {
	node := n
	tmpitem := new(Item)
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

func (n *Node) isExist(b byte) *Item {
	for _, v := range n.items {
		if v.val == b {
			return v
		}
	}
	return nil
}
