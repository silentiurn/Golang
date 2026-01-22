package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	current := l
	for current != nil && pos != count {
		current = current.Next
		count++
	}
	return current
}
