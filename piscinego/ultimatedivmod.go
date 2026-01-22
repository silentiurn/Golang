package piscine

func UltimateDivMod(a *int, b *int) {
	if a == nil || b == nil || *b == 0 {
		return
	}
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}
