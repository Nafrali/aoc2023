package main

import "sort"

type interval struct {
	start int
	end   int
}

type intervals struct {
	intervals []*interval
}

func RemoveIndex(s []*interval, index int) []*interval {
	return append(s[:index], s[index+1:]...)
}

func (i *intervals) addInterval(interval *interval) {
	startBefore, index := i.getIntervalAtOrBeforeStart(interval.start)
	i.mergeAndDeleteOldInterval(startBefore, index, interval)

	startAfter, index := i.getIntervalAtOrAfterEnd(interval.start)
	i.mergeAndDeleteOldInterval(startAfter, index, interval)

	startMiddle := i.getIntervalsDuringInterval(interval)
	sort.Slice(startMiddle, func(a, b int) bool {
		return startMiddle[b] < startMiddle[a]
	})
	for _, index = range startMiddle {
		RemoveIndex(i.intervals, index)
	}
	i.intervals = append(i.intervals, interval)
}

func (i *intervals) mergeAndDeleteOldInterval(oldInterval *interval, index int, newInterval *interval) {
	if oldInterval == nil {
		return
	}
	newInterval.start = min(oldInterval.start, newInterval.start)
	newInterval.end = max(oldInterval.end, newInterval.end)
	i.intervals = RemoveIndex(i.intervals, index)
}

func (i *intervals) getIntervalsDuringInterval(newInterval *interval) []int {
	var indexes = []int{}

	for index, t := range i.intervals {
		if t.end < newInterval.end && t.start > newInterval.start {
			indexes = append(indexes, index)
			newInterval.start = min(newInterval.start, t.start)
			newInterval.end = max(newInterval.end, t.end)
		}
	}
	return indexes
}

func (i *intervals) getIntervalAtOrAfterEnd(limit int) (*interval, int) {
	for index, t := range i.intervals {
		if t.end == limit || t.end == limit+1 {
			return t, index
		}
	}
	return nil, -1
}

func (i *intervals) getIntervalAtOrBeforeStart(limit int) (*interval, int) {
	for index, t := range i.intervals {
		if t.start == limit || t.start == limit-1 {
			return t, index
		}
	}
	return nil, -1
}
