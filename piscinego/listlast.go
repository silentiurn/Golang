package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil { // защита от пустого списка
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
