package algo

// PascalTri -- simply calculate one block
func PascalTri(i, j int) int {
	if !(j == 1 || i == j) {
		return PascalTri(i-1, j-1) + PascalTri(i-1, j)
	}
	return 1
}

// GeneratePascalTri -- generate pascal triangle 0 to numRows
func GeneratePascalTri(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := range tmp {
			tmp[j] = PascalTri(i+1, j+1)
		}
		result[i] = tmp
	}
	return result
}

//GetRowPascalTri -- get one row of PascalTri
func GetRowPascalTri(rowIndex int) []int {
	ReArray := make([]int, rowIndex+1)
	ReArray[rowIndex], ReArray[0] = 0, 1
	for i := 0; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			ReArray[j] += ReArray[j-1]
		}
	}
	return ReArray
}
