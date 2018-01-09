package pythagorean

type Triplet [3]int

func Range(min, max int) (res []Triplet) {
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i+j*j == k*k {
					res = append(res, Triplet{i, j, k})
				}
			}
		}
	}
	return
}

func Sum(p int) (res []Triplet) {
	for i := 1; i <= p/3; i++ {
		for j := i; j <= p/2; j++ {
			k := p - i - j
			if i*i+j*j == k*k {
				res = append(res, Triplet{i, j, k})
			}
		}
	}
	return

}
