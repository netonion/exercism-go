package pascal

// Triangle produces a pascal triangle of n rows.
func Triangle(n int) (res [][]int) {
	res = append(res, []int{1})
	for i := 1; i < n; i++ {
		row := []int{1}
		for j := 1; j < i; j++ {
			row = append(row, res[i-1][j-1]+res[i-1][j])
		}
		row = append(row, 1)
		res = append(res, row)
	}
	return
}
