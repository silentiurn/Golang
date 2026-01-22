package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	var prev *NodeL
	for current != nil {
		if current.Data == data_ref {
			if current == l.Head {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
		}
		if current.Data != data_ref {
			prev = current
		}
		current = current.Next
	}
}
