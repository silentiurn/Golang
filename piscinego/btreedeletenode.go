package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
	} else {
		// Находим минимальный элемент в правом поддереве (successor)
		successor := BTreeMin(node.Right)
		if successor.Parent != node {
			root = BTreeTransplant(root, successor, successor.Right)
			successor.Right = node.Right
			if successor.Right != nil {
				successor.Right.Parent = successor
			}
		}
		root = BTreeTransplant(root, node, successor)
		successor.Left = node.Left
		if successor.Left != nil {
			successor.Left.Parent = successor
		}
	}

	// Очистка ссылок удаляемого узла
	node.Left = nil
	node.Right = nil
	node.Parent = nil

	return root
}
