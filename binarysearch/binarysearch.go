package binarysearch

func BinarySearchAsc(arr []int, target int) int {
	pointer := len(arr) / 2
	var idxAccessed map[int]bool = make(map[int]bool)
	for !idxAccessed[pointer] {
		idxAccessed[pointer] = true
		if arr[pointer] == target {
			return pointer
		}
		if arr[pointer] > target {
			pointer = pointer / 2
		}
		if arr[pointer] < target {
			pointer = pointer + ((len(arr) - pointer) / 2)
		}
	}
	return -1
}

func BinarySearchDesc(arr []int, target int) int {
	pointer := len(arr) / 2
	var idxAccessed map[int]bool = make(map[int]bool)
	for !idxAccessed[pointer] {
		idxAccessed[pointer] = true
		if arr[pointer] == target {
			return pointer
		}
		if arr[pointer] < target {
			pointer = pointer / 2
		}
		if arr[pointer] > target {
			pointer = pointer + ((len(arr) - pointer) / 2)
		}
	}
	return -1
}
