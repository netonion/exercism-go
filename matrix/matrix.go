package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type matrix [][]int

func New(s string) (matrix, error) {
	var m matrix
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Fields(line)
		if len(m) > 0 && len(fields) != len(m[len(m)-1]) {
			return nil, errors.New("wrong number of fields")
		}
		row := make([]int, len(fields))
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row[i] = num
		}
		m = append(m, row)
	}
	return m, nil
}

func (m matrix) Rows() [][]int {
	n := make([][]int, len(m))
	for i, row := range m {
		n[i] = make([]int, len(row))
		copy(n[i], row)
	}
	return n
}

func (m matrix) Cols() [][]int {
	if len(m) == 0 {
		return [][]int{}
	}
	n := make([][]int, len(m[0]))
	for _, row := range m {
		for i, val := range row {
			n[i] = append(n[i], val)
		}
	}
	return n
}

func (m matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(m) || (len(m) > 0 && c >= len(m[0])) {
		return false
	}
	m[r][c] = val
	return true
}
