package nc0022

func merge(A []int, m int, B []int, n int) {
	// write code here

	i, j := m-1, n-1
	for k := m + n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 {
			if A[i] > B[j] {
				A[k] = A[i]
				i--
			} else {
				A[k] = B[j]
				j--
			}
		} else if i >= 0 {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
	}
}
