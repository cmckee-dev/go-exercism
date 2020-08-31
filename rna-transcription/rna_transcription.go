package strand

var transcribe = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var rna []rune
	for _, nucleotide := range dna {
		rna = append(rna, transcribe[nucleotide])
	}

	return string(rna)
}
