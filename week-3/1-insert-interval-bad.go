func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 1 {
		retEarly := false
		// if single intervals is to be merged
		if (newInterval[1] >= intervals[0][1] && newInterval[0] <= intervals[0][1]) || (newInterval[0] <= intervals[0][0] && newInterval[1] >= intervals[0][0]) {
			retEarly = true
			if newInterval[1] > intervals[0][1] {
				intervals[0][1] = newInterval[1]
			}
			if newInterval[0] < intervals[0][0] {
				intervals[0][0] = newInterval[0]
			}

			// if newInterval is within single interval
		} else if newInterval[0] >= intervals[0][0] && newInterval[0] <= intervals[0][1] && newInterval[1] <= intervals[0][1] {
			retEarly = true
		}
		if retEarly {
			return intervals
		}
	}

	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	mergeStart, mergeEnd, mergeBeginIdx, mergeEndIdx := 0, 0, -1, len(intervals)-1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] || intervals[i][1] <= intervals[i-1][1] || intervals[i][0] <= mergeEnd {
			// merge
			if mergeBeginIdx == -1 {
				mergeStart = intervals[i-1][0]
				mergeEnd = intervals[i-1][1]
			} else if intervals[i-1][0] < mergeStart {
				mergeStart = intervals[i-1][0]
			}

			if mergeBeginIdx == -1 {
				mergeBeginIdx = i - 1
			}

			if intervals[i][1] > mergeEnd {
				mergeEnd = intervals[i][1]
			}

		} else if mergeBeginIdx != -1 {
			mergeEndIdx = i - 1
			break
		}

	}

	if mergeBeginIdx != -1 {

		if mergeBeginIdx == 0 {
			intervals[mergeEndIdx][0] = mergeStart
			intervals[mergeEndIdx][1] = mergeEnd
			intervals = intervals[mergeEndIdx:]

		} else if mergeBeginIdx == len(intervals)-1 {
			intervals[mergeBeginIdx][0] = mergeStart
			intervals[mergeBeginIdx][1] = mergeEnd
			intervals = intervals[0 : mergeBeginIdx+1]
		} else {
			intervals[mergeBeginIdx][0] = mergeStart
			intervals[mergeBeginIdx][1] = mergeEnd
			left, right := mergeBeginIdx+1, mergeEndIdx+1
			for right < len(intervals) {
				intervals[left], intervals[right] = intervals[right], intervals[left]
				left++
				right++
			}
			intervals = intervals[0:left]
		}
	}
	return intervals
}


