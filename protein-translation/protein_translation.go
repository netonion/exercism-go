package protein

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon translates a codon to protein
func FromCodon(codon string) string {
	return codonToProtein[codon]
}

// FromRNA translates a RNA sequence to proteins
func FromRNA(rna string) (res []string) {
	for i := 0; i < len(rna); i += 3 {
		if protein := FromCodon(rna[i : i+3]); protein == "STOP" {
			break
		} else {
			res = append(res, protein)
		}
	}
	return
}
