package time_based_key_value_store

type TimeStampedValue struct {
	value     string
	timestamp int
}

type TimeMap struct {
	timeStampedValuesByKey map[string][]TimeStampedValue
}

func Constructor() TimeMap {
	return TimeMap{timeStampedValuesByKey: make(map[string][]TimeStampedValue)}
}

func (timeMap *TimeMap) Set(key string, value string, timestamp int) {
	timeStampedValue := TimeStampedValue{value: value, timestamp: timestamp}
	values, ok := timeMap.timeStampedValuesByKey[key]
	if !ok {
		timeMap.timeStampedValuesByKey[key] = append([]TimeStampedValue{}, timeStampedValue)
	} else {
		index := findInsertIndexIn(values, timeStampedValue)
		values = append(values, TimeStampedValue{})
		copy(values[index+1:], values[index:])
		values[index] = timeStampedValue
		timeMap.timeStampedValuesByKey[key] = values
	}
}

func (timeMap *TimeMap) Get(key string, timestamp int) string {
	values, ok := timeMap.timeStampedValuesByKey[key]
	if !ok {
		return ""
	}
	return getNearestIn(values, timestamp)
}

func getNearestIn(values []TimeStampedValue, timestamp int) string {
	var getNearestInner func(start, end, offset int) string
	getNearestInner = func(start, end, offset int) string {
		trimmedSource := values[start:end]
		if len(trimmedSource) == 0 {
			return ""
		}
		mid := len(trimmedSource)/2 + offset
		if timestamp == values[mid].timestamp {
			return values[mid].value
		} else if timestamp < values[mid].timestamp {
			if mid-1 >= 0 {
				if values[mid-1].timestamp < timestamp {
					return values[mid-1].value
				}
				return getNearestInner(start, mid, 0)
			}
			return ""
		} else {
			if mid+1 < len(values) {
				if values[mid+1].timestamp > timestamp {
					return values[mid].value
				}
				return getNearestInner(mid+1, end, mid+1)
			}
			if values[mid].timestamp < timestamp {
				return values[mid].value
			}
			return ""
		}
	}
	return getNearestInner(0, len(values), 0)
}

func findInsertIndexIn(values []TimeStampedValue, value TimeStampedValue) int {
	var findIndexInner func(start, end, offset int) int
	findIndexInner = func(start, end, offset int) int {
		trimmedSource := values[start:end]
		if len(trimmedSource) == 0 {
			return start
		}
		mid := len(trimmedSource)/2 + offset
		if value.timestamp <= values[mid].timestamp {
			return findIndexInner(start, mid, 0)
		} else {
			return findIndexInner(mid+1, end, mid+1)
		}
	}
	return findIndexInner(0, len(values), 0)
}
