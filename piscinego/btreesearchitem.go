package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	if left := BTreeSearchItem(root.Left, elem); left != nil {
		return left
	}
	return BTreeSearchItem(root.Right, elem)
}
