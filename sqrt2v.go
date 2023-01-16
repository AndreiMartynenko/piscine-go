package piscine

func Sqrt2(nb int) int {
	if nb == 1 {
		return 1
	}

	if nb == 0 {
		return 0
	}

	for res := 1; res < nb; res++ {
		mod := nb % res
		if mod == 0 {
			div := nb / res
			if div*div == nb {
				return res
			}
		}

	}
	return 0
}
