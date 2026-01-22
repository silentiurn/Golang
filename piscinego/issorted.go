package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	// Проверяем направление сортировки по первым двум различным элементам
	direction := 0
	for i := 0; i < len(a)-1 && direction == 0; i++ {
		cmp := f(a[i], a[i+1])
		if cmp > 0 {
			direction = -1 // убывающая
		} else if cmp < 0 {
			direction = 1 // возрастающая
		}
		// если cmp == 0, продолжаем искать направление
	}

	// Если все элементы равны или слайс из одного элемента
	if direction == 0 {
		return true
	}

	// Проверяем сохранение направления по всему слайсу
	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])
		if direction > 0 && cmp > 0 { // для возрастающей: не должно быть a[i] > a[i+1]
			return false
		}
		if direction < 0 && cmp < 0 { // для убывающей: не должно быть a[i] < a[i+1]
			return false
		}
	}

	return true
}
