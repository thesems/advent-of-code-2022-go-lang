package utils

// https://github.com/google/guava/blob/master/guava/src/com/google/common/collect/EvictingQueue.java
type EvictingQueue struct {
	data []int
	size int
}

func CreateEvictingQueue(size int) *EvictingQueue {
	data := []int{}
	queue := &EvictingQueue{data, size}
	return queue
}

func (q *EvictingQueue) Distinct() bool {
	// https://www.youtube.com/watch?v=YQs6IC-vgmo&t=0s&ab_channel=AlessandroStamatto
	// Method 1. use maps to detect duplicates (via collisions)
	// m := make(map[int]bool)
	// for i := 0; i < len(q.data); i++ {
	// 	_, ok := m[i]
	// 	if ok {
	// 		return false
	// 	}
	// }
	// return true

	// Method 2. use linear search with worst-case O(n^2)
	for i := 0; i < len(q.data); i++ {
		for j := 0; j < len(q.data); j++ {
			if q.data[i] == q.data[j] && i != j {
				return false
			}
		}
	}
	return true
}

func (q *EvictingQueue) Add(value int) {
	if len(q.data) != q.size {
		q.data = append(q.data, value)
	} else {
		q.data = q.data[1:]
		q.data = append(q.data, value)
	}
}
