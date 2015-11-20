package algorithms

type NeeWun struct {
	gap            int
	missmatchscore int
	exactscore     int
	StringDistance
}

func NewNeeWun(gap int, missmatchscore int, exactscore int) *NeeWun {
	return &NeeWun{gap, missmatchscore, exactscore, nil}
}

func (dist NeeWun) CalculateDistance(fromString string, toString string) int {
	lenFromString := len(fromString)
	lenToString := len(toString)

	matrix := make([][]int, lenFromString+1)
	matrixroute := make([][]int, lenFromString+1)
	for i := 0; i < lenFromString+1; i++ {
		matrix[i] = make([]int, lenToString+1)
		matrixroute[i] = make([]int, lenToString+1)
	}

	matrix[0][0] = 0
	for j := 1; j < lenToString+1; j++ {
		matrix[0][j] = matrix[0][j-1] + dist.gap
		matrixroute[0][j] = int(Left)
	}

	for i := 1; i < lenFromString+1; i++ {
		matrix[i][0] = matrix[i-1][0] + dist.gap
		matrixroute[i][0] = int(Up)

		for j := 1; j < lenToString+1; j++ {
			var score int
			if []byte(toString)[j-1] == []byte(fromString)[i-1] {
				score = dist.exactscore
			} else {
				score = -1 * dist.missmatchscore
			}

			gdiag := matrix[i-1][j-1] + score
			gup := matrix[i-1][j] + dist.gap
			gleft := matrix[i][j-1] + dist.gap

			result, route := maximumwithdirection3(gdiag, gup, gleft)
			matrix[i][j] = result
			matrixroute[i][j] = int(route)
		}

	}

	logArrayLine(matrix)
	logArrayLine(matrixroute)

	distance, _ := traceback(matrixroute)
	return distance
}

func traceback(matrixroute [][]int) (int, int) {
	i := len(matrixroute) - 1
	j := len(matrixroute[i]) - 1
	walk := 0
	distance := 0
	for !(i == 0 && j == 0) {
		currentarrow := matrixroute[i][j]
		switch currentarrow {
		case int(Diag):
			i--
			j--
		case int(Left):
			j--
			distance++
		case int(Up):
			i--
			distance++
		}
		walk++
	}
	return distance, walk
}