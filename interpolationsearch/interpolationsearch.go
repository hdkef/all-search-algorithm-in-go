package interpolationsearch

func InterpolationSearchAsc(arr []int, target int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + ((target-arr[low])/(arr[high]-arr[low]))*(high-low)
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

func InterpolationSearchDesc(arr []int, target int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + ((target-arr[low])/(arr[high]-arr[low]))*(high-low)
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
