package main

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func rotateRunes(s string, shift int) string {
	if s == "" {
		return s
	}
	
	left := shift > 0
	rns := []rune(s)
	n := len(rns)
	ans := make([]rune, n)
	shift = abs(shift % n)
	
	if shift == 0 {
		return string(rns)
	}
	
	for i := range rns {
		newI := 0
		if left {
			newI = (i - shift + n) % n
		} else {
			newI = (i + shift) % n
		}

		ans[newI] = rns[i]
	}

	return string(ans)
}
