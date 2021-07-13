package mutants

type Stats struct {
	MutantDNACount int64   `json:"count_mutant_dna"`
	HumanDNACount  int64   `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}
