package mutants

const (
	minSequenceForMutant = 1
	minLetterForSequence = 3
)

type fn func(dna []string, x, i, c int) bool

var fns = []fn{horizontally, down, obliqueRight, obliqueLeft}

func isMutant(dna []string) bool {
	var c int
	for x := range dna {
		for i := 0; i < len(dna[x]); i++ {
			for _, f := range fns {
				if f(dna, x, i, 0) {
					c++
					if c > minSequenceForMutant {
						return true
					}
				}
			}
		}
	}
	return false
}

func horizontally(dna []string, x, i, c int) bool {
	if c >= minLetterForSequence {
		return true
	}
	if i < len(dna[x])-1 &&
		len(dna[x])-i > minLetterForSequence-c {
		if dna[x][i] == dna[x][i+1] {
			return horizontally(dna, x, i+1, c+1)
		} else {
			return false
		}
	}
	return false
}

func down(dna []string, x, i, c int) bool {
	if c >= minLetterForSequence {
		return true
	}
	if x < len(dna)-1 &&
		len(dna)-x > minLetterForSequence-c {
		if dna[x][i] == dna[x+1][i] {
			return down(dna, x+1, i, c+1)
		} else {
			return false
		}
	}
	return false
}

func obliqueRight(dna []string, x, i, c int) bool {
	if c >= minLetterForSequence {
		return true
	}
	if x < len(dna)-1 && i < len(dna[x])-1 &&
		len(dna)-x > minLetterForSequence-c && len(dna[x])-i > minLetterForSequence-c {
		if dna[x][i] == dna[x+1][i+1] {
			return obliqueRight(dna, x+1, i+1, c+1)
		} else {
			return false
		}
	}
	return false
}

func obliqueLeft(dna []string, x, i, c int) bool {
	if c >= minLetterForSequence {
		return true
	}
	if x < len(dna)-1 && i > 0 &&
		len(dna)-x > minLetterForSequence-c && i >= minLetterForSequence-c {
		if dna[x][i] == dna[x+1][i-1] {
			return obliqueLeft(dna, x+1, i-1, c+1)
		} else {
			return false
		}
	}
	return false
}
