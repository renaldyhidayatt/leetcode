package firstbadversion

type VersionControl struct {
	firstBadVersion int
}

func (v *VersionControl) isBadVersion(version int) bool {
	if version >= v.firstBadVersion {
		return true
	}
	return false
}

type First struct {
	VersionControl
}

func (f *First) firstBadVersion(n int) int {
	if n < 1 {
		return 0
	}
	start := 1
	end := n
	for start+1 < end {
		mid := start + (end-start)/2
		if f.VersionControl.isBadVersion(mid) {
			end = mid
		} else {
			start = mid
		}
	}

	if f.VersionControl.isBadVersion(start) {
		return start
	}
	if f.VersionControl.isBadVersion(end) {
		return end
	}
	return 0
}
