package sort

func hoareQuickSort(arr []int, lo int, hi int) {
	if lo < hi {
		p := hoarePartition(arr, lo, hi)
		hoareQuickSort(arr, lo, p)
		hoareQuickSort(arr, p + 1, hi)
	}
}

func lomutoQuickSort(arr []int, lo int, hi int) {
	if lo < hi {
		p := lomutoPartition(arr, lo, hi)
		lomutoQuickSort(arr, lo, p - 1)
		lomutoQuickSort(arr, p + 1, hi)
	}
}

// A quicksort implementation using the Lomuto lomutoPartition scheme.
// The pivot is chosen as the last element of the array, and the pivot is moved to its
// final position at the end of the iteration, resulting in all elements in the array before
// the pivot being lower than the pivot and all elements after being greater.
func LomutoQuickSort(arr []int) {
 	lomutoQuickSort(arr, 0, len(arr) - 1)
}

func lomutoPartition(arr []int, lo int, hi int) int {
	var pivot = arr[hi]
	var i = lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[hi], arr[i+1] = arr[i + 1], arr[hi]
	return i + 1
}

func HoareQuickSort(arr []int) {
	hoareQuickSort(arr, 0, len(arr) - 1)
}

func hoarePartition(arr []int, lo int, hi int) int {
	var pivot = arr[lo + (hi - lo) / 2]
	var i = lo - 1
	var j = hi + 1
	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}
