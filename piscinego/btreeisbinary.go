package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBinaryHelper(root, "", "")
}

func isBinaryHelper(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}

	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}

	return isBinaryHelper(node.Left, min, node.Data) &&
		isBinaryHelper(node.Right, node.Data, max)
}
