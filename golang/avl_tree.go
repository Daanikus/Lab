package golang

type Node struct {
	val    int
	height int
	left   *Node
	right  *Node
}

func GetBalanceFactor(node * Node) int {
	if node != nil {
		return node.left.height - node.right.height
	} else {
		return 0
	}

}

type AVLTree struct {
	root *Node
}


func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	} else {
		return max(getHeight(node.left), getHeight(node.right)) + 1
	}
}

func leftRotate(cur *Node) *Node {
	left := cur.left
	cur.left = left.right
	left.right = cur

	// update height info
	left.height = getHeight(left)
	cur.height = getHeight(cur)
	return left
}

func rightRotate(cur *Node) *Node {
	right := cur.right
	cur.right = right.left
	right.left = cur

	// update height info
	right.height = getHeight(right)
	cur.height = getHeight(cur)
	return right
}

func Rebalance(root *Node) *Node {
	fac := GetBalanceFactor(root)
	if fac > 1 && GetBalanceFactor(root.left) > 0 {  // LL
		return rightRotate(root)
	} else if fac > 1 && GetBalanceFactor(root.left) < 0 {  // LR
		root.left = leftRotate(root.left)
		return rightRotate(root)
	} else if fac < -1 && GetBalanceFactor(root.right) < 0 {  // RR
		return leftRotate(root)
	} else if fac < -1 && GetBalanceFactor(root.right) > 0 {  // RL
		root.right = rightRotate(root.right)
		return leftRotate(root)
	} else {
		return root
	}
}

func Insert(cur *Node, val int) {
	if cur == nil {
		cur = new(Node)
		cur.val = val
	} else if cur.val == val {
		return
	} else {
		if val < cur.val {
			Insert(cur.left, val)
		} else {
			Insert(cur.right, val)
		}
	}

	Rebalance(cur)
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}
