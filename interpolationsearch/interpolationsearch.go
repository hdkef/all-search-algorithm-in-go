package interpolationsearch

func InterpolationSearchAsc(arr []int, target int) int {
	mid := len(arr) / 2
	low := 0
	high := len(arr) - 1

	for low <= high && arr[mid] >= arr[low] && arr[mid] <= arr[high] {
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			mid = low + ((arr[mid]-arr[low])/(arr[high]-arr[low]))*(high-low)
			low = mid + 1
		}
		if arr[mid] > target {
			mid = low + ((arr[mid]-arr[low])/(arr[high]-arr[low]))*(high-low)
			high = mid - 1
		}
	}

	return -1
}
