package mutants

func isMutant(dna []string) bool {
	for x := range dna {
		for i := 0; i < len(dna[x]); i++ {
			if horizontally(dna, x, i, 0) ||
				down(dna, x, i, 0) ||
				obliqueRight(dna, x, i, 0) ||
				obliqueLeft(dna, x, i, 0) {
				return true
			}
		}
	}
	return false
}

func horizontally(dna []string, x, i, c int) bool {
	if c >= 3 {
		return true
	}
	if i < len(dna[x])-1 {
		if dna[x][i] == dna[x][i+1] {
			return horizontally(dna, x, i+1, c+1)
		} else {
			return false
		}
	}
	return false
}

func down(dna []string, x, i, c int) bool {
	if c >= 3 {
		return true
	}
	if x < len(dna)-1 {
		if dna[x][i] == dna[x+1][i] {
			return down(dna, x+1, i, c+1)
		} else {
			return false
		}
	}
	return false
}

func obliqueRight(dna []string, x, i, c int) bool {
	if c >= 3 {
		return true
	}
	if x < len(dna)-1 && i < len(dna[x])-1 {
		if dna[x][i] == dna[x+1][i+1] {
			return obliqueRight(dna, x+1, i+1, c+1)
		} else {
			return false
		}
	}
	return false
}

func obliqueLeft(dna []string, x, i, c int) bool {
	if c >= 3 {
		return true
	}
	if x < len(dna)-1 && i > 0 {
		if dna[x][i] == dna[x+1][i-1] {
			return obliqueLeft(dna, x+1, i-1, c+1)
		} else {
			return false
		}
	}
	return false
}
