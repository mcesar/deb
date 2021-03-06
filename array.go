package deb

import (
	_ "fmt"
)

type Array [][][]int64

func NewArray(x, y, z int) (result Array) {
	result = make(Array, x)
	values := make([]int64, x*y*z)
	for i := range result {
		result[i] = make([][]int64, y)
		for j := range result[i] {
			result[i][j], values = values[:z], values[z:]
		}
	}
	return
}

func (arr Array) Copy() (result Array) {
	if arr.Empty() {
		return Array{{{}}}
	}
	x, y, z := arr.Dimensions()
	result = NewArray(x, y, z)
	for i := range result {
		for j := range result[i] {
			copy(result[i][j], arr[i][j])
		}
	}
	return
}

func (arr Array) Transposed() (result Array) {
	if arr.Empty() {
		return Array{{{}}}
	}
	x, y, z := arr.Dimensions()
	result = NewArray(z, y, x)
	for i := range result {
		for j := range result[i] {
			for k := range result[i][j] {
				result[i][j][k] = arr[k][j][i]
			}
		}
	}
	return
}

func (arr *Array) Append(other_arr *Array, y, z uint64) error {
	x1, y1, z1 := arr.Dimensions()
	x2, y2, z2 := other_arr.Dimensions()
	mx, my, mz := max(x1, x2), maxUint64(uint64(y1), uint64(y2)+y),
		maxUint64(uint64(z1), uint64(z2)+z)
	if arr.Empty() && other_arr.Empty() {
		return nil
	}
	result := make(Array, mx)
	values := make([]int64, uint64(mx)*my*mz)
	for i := range result {
		result[i] = make([][]int64, my)
		for j := range result[i] {
			result[i][j], values = values[:mz], values[mz:]
			for k := range result[i][j] {
				if i < x1 && j < y1 && k < z1 {
					result[i][j][k] = (*arr)[i][j][k]
				}
				if i < x2 && j < y2 && k < z2 {
					result[i][j][uint64(k)+z] = (*other_arr)[i][j][k]
				}
			}
		}
	}
	*arr = result
	return nil
}

func (arr Array) Dimensions() (int, int, int) {
	if len(arr) == 0 || len(arr[0]) == 0 || len(arr[0][0]) == 0 {
		return 0, 0, 0
	} else {
		return len(arr), len(arr[0]), len(arr[0][0])
	}
}

func (arr Array) Empty() bool {
	x, y, z := arr.Dimensions()
	return x == 0 || y == 0 || z == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
