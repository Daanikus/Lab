package golang

type KeyType int
type PrimaryKeyType int

type KV struct {
	key        KeyType
	primaryKey PrimaryKeyType
}

type MiddleNode struct {
	cap   int
	par *MiddleNode
	nodes []KeyType
	middleLinks []*MiddleNode
	leafLinks []*LeafNode
}

func (m *MiddleNode) isEmpty() bool {
	return len(m.nodes) > 0
}

func (m *MiddleNode) isFull() bool {
	return m.cap == len(m.nodes)
}

func (m *MiddleNode) insert() {

}

func (m *MiddleNode) split() (newRightLeafNode *MiddleNode) {

}

type LeafNode struct {
	cap int
	nodes []KV
	par *MiddleNode
	left  *LeafNode
	right *LeafNode
}

func (m *LeafNode) isEmpty() bool {
	return len(m.nodes) > 0
}

func (m *LeafNode) isFull() bool {
	return m.cap == len(m.nodes)
}

func (m *LeafNode) insert() {

}

func (m *LeafNode) split() (newRightLeafNode *LeafNode) {

}

type LeafNodeIterator interface {
	Next() *LeafNode
}

type BPTree struct {
	root     *MiddleNode
	leafHead *LeafNode
}

type bPTree interface {
	getLeaves() *LeafNodeIterator
	Get(key KeyType) PrimaryKeyType
	Set(key KeyType, prim PrimaryKeyType)
	Scan(lowKey KeyType, highKey KeyType, lowKeyInclusive bool, highKeyInclusive bool) *LeafNodeIterator
}

func seek(key KeyType, curNode *MiddleNode, leafHead *LeafNode) *LeafNode {
	retPoint := -1
	if !curNode.isEmpty() {
		for i, v := range curNode.nodes {
			if key < v {
				retPoint = i
			}
		}
		if retPoint == -1 {
			retPoint = curNode.cap
		}

		if len(curNode.middleLinks) > 0 {
			return seek(key, curNode.middleLinks[retPoint], leafHead)
		} else {
			return curNode.leafLinks[retPoint]
		}
	} else {
		return leafHead
	}
}

func (b *BPTree) getLeaves() *LeafNodeIterator {
	panic("implement me")
}

func (b *BPTree) Get(key KeyType) PrimaryKeyType {
	panic("implement me")
}

func (b *BPTree) Set(key KeyType, prim PrimaryKeyType) {
	panic("implement me")
}

func (b *BPTree) Scan(lowKey KeyType, highKey KeyType, lowKeyInclusive bool, highKeyInclusive bool) *LeafNodeIterator {
	panic("implement me")
}

func NewMiddleNode(middleNodeCap int) *MiddleNode {
	return &MiddleNode{
		middleNodeCap,
		nil,
		make([]KeyType, middleNodeCap + 1),
		make([]*MiddleNode, middleNodeCap + 1),
		make([]*LeafNode, middleNodeCap + 1),
	}
}

func NewLeafNode(middleNodeCap int) *LeafNode {
	return &LeafNode{
		cap: middleNodeCap,
		nodes: make([]KV, middleNodeCap + 1),
		par: nil,
		left:  nil,
		right: nil,
	}
}

func NewBPTree(middleNodeCap int) *BPTree {
	return &BPTree{
		root:     NewMiddleNode(middleNodeCap),
		leafHead: NewLeafNode(middleNodeCap),
	}
}
