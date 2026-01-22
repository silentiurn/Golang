package piscine

func ListReverse(l *List) {
	var prev *NodeL   // предыдущий элемент
	current := l.Head // текущий элемент
	var next *NodeL   // временный для хранения следующего

	// Меняем направления связей
	for current != nil {
		next = current.Next // запоминаем следующий
		current.Next = prev // переворачиваем ссылку
		prev = current      // двигаем prev вперёд
		current = next      // двигаем current вперёд
	}

	// Меняем местами Head и Tail
	l.Tail = l.Head
	l.Head = prev
}
