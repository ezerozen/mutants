package mutants

import "testing"

func Test_IsMutant(t *testing.T) {
	type args struct {
		dna []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is mutant",
			args: args{dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}},
			want: true,
		},
		{
			name: "is mutant",
			args: args{dna: []string{"ATGCGA", "CAGTAC", "TTAATT", "AGACGG", "GCGTCA", "TCACTG"}},
			want: true,
		},
		{
			name: "is mutant",
			args: args{dna: []string{"ATGCGAATC", "CAGTACGCT", "TTAATTCTA", "AGACGGAGT", "GCGTCAGAT", "TCACTGCGA",
				"AATCGCTAA", "GGCTAACGT", "AATTCCGGA"}},
			want: true,
		},
		{
			name: "is human",
			args: args{dna: []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMutant(tt.args.dna); got != tt.want {
				t.Errorf("IsMutant() = %v, want %v", got, tt.want)
			}
		})
	}
}
