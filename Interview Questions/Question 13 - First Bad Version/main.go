package main

//
func isBadVersion(n int) bool {
	return true
}

func firstBadVersion(n int) int {
	if isBadVersion(1) {
		return 1
	}

	if n == 2 {
		if isBadVersion(2) {
			return 2
		}
	}

	start := 1
	end := n
	mid := (end - start) / 2

	for end-start > 1 {
		if isBadVersion(mid) {
			end = mid
			mid = (end - start) / 2
		} else {
			start = mid
			mid = mid + (end-start)/2
		}
	}
	return end
}

func main() {

}
