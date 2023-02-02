package binarysearch

func BinarySearchAsc(arr []int, target int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := high + low/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		}
		if arr[mid] < target {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchDesc(arr []int, target int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := high + low/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			high = mid - 1
		}
		if arr[mid] > target {
			low = mid + 1
		}
	}
	return -1
}
