package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l2 == nil {
		return
	}
	// если первый список пустой — просто присваиваем второй
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// если второй список пустой — ничего не делаем
	if l2.Head == nil {
		return
	}

	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
