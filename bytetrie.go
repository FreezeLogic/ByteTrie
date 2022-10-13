package bytetrie

type Node struct {
	Items []*Item
}
type Item struct {
	Val   byte
	Last  bool
	Child *Node
}

func NewTrie() *Node {
	n := new(Node)
	n.Items = *new([]*Item)
	return n
}

func (n *Node) Insert(b []byte) {
	node := n
	tmpitem := new(Item)
	for _, vb := range b {
		newitem := node.isExist(vb)
		if newitem == nil {
			node.Items = append(node.Items, &Item{
				Val:   vb,
				Last:  false,
				Child: NewTrie(),
			})
			tmpitem = node.Items[len(node.Items)-1]
			node = tmpitem.Child
		} else {
			tmpitem = newitem
			node = tmpitem.Child
		}
	}
	tmpitem.Last = true
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
			node = tmpitem.Child
		}
	}
	if tmpitem.Last {
		return true
	}
	return false
}

func (n *Node) isExist(b byte) *Item {
	for _, v := range n.Items {
		if v.Val == b {
			return v
		}
	}
	return nil
}
