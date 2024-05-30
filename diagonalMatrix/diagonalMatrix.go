package diagonalmatrix

func DiagonalDifference(arr [][]int32) int32 {
	var num1 int32
	var num2 int32
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		num1 += arr[i][i]
		num2 += arr[i][arrLen-i-1]
	}
	num1 = num1 - num2
	if num1 < 0 {
		num1 *= -1

	}

	return num1
}
