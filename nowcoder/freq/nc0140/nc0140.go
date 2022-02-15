package nc0140

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort(arr []int) []int {
	// write code here
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func MySort_Bubble(arr []int) []int {
	// write code here
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func MySort_Quick(arr []int) []int {
	// write code here
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIdx := partition(arr, low, high)
		quickSort(arr, low, pivotIdx-1)
		quickSort(arr, pivotIdx+1, high)
	}
}
