package snapshot_array

import (
	"sort"
)

type SnapshotArray struct {
	snapshotId int
	elements   [][]snapshotValue
}

type snapshotValue struct {
	value      int
	snapshotId int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapshotId: 0,
		elements:   make([][]snapshotValue, length),
	}
}

func (snapshotArray *SnapshotArray) Set(index int, val int) {
	snapshots := snapshotArray.elements[index]
	if len(snapshots) == 0 {
		snapshots = append(snapshots, snapshotValue{snapshotId: snapshotArray.snapshotId, value: val})
		snapshotArray.elements[index] = snapshots
	} else {
		position := snapshotArray.search(snapshots, snapshotArray.snapshotId)
		if position < len(snapshots) && snapshots[position].snapshotId == snapshotArray.snapshotId {
			snapshots[position] = snapshotValue{snapshotId: snapshotArray.snapshotId, value: val}
		} else {
			snapshots = append(snapshots, snapshotValue{snapshotId: snapshotArray.snapshotId, value: val})
			snapshotArray.elements[index] = snapshots
		}
	}
}

func (snapshotArray *SnapshotArray) Snap() int {
	id := snapshotArray.snapshotId
	snapshotArray.snapshotId = snapshotArray.snapshotId + 1
	return id
}

func (snapshotArray *SnapshotArray) Get(index int, snapShotId int) int {
	snapshots := snapshotArray.elements[index]
	if len(snapshots) == 0 {
		return 0
	}
	if snapShotId > snapshots[len(snapshots)-1].snapshotId {
		return snapshots[len(snapshots)-1].value
	}
	position := snapshotArray.search(snapshots, snapShotId)
	if position < len(snapshots) && snapshots[position].snapshotId == snapShotId {
		return snapshots[position].value
	}
	if snapShotId < snapshots[position].snapshotId && position > 0 {
		return snapshots[position-1].value
	}
	if snapShotId > snapshots[position].snapshotId {
		return snapshots[position].value
	}
	return 0
}

func (snapshotArray *SnapshotArray) search(snapshots []snapshotValue, snapshotId int) int {
	return sort.Search(len(snapshots), func(index int) bool {
		if snapshots[index].snapshotId >= snapshotId {
			return true
		}
		return false
	})
}
