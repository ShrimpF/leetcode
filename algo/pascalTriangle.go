package algo

func pascalTri(i, j int) int {
	if !(j == 1 || i == j) {
		return pascalTri(i-1, j-1) + pascalTri(i-1, j)
	}
	return 1
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := range tmp {
			tmp[j] = pascalTri(i+1, j+1)
		}
		result[i] = tmp
	}
	return result
}

func getRow(rowIndex int) []int {
	ReArray := make([]int, rowIndex+1)
	ReArray[rowIndex], ReArray[0] = 0, 1
	for i := 0; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			ReArray[j] += ReArray[j-1]
		}
	}
	return ReArray
}
