package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	del := 0
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			del = del + 1
			if del == 1 {
				return false
			}
		}
	}
	return true
}
