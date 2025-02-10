package sort

func BucketSort(array []int) []int {
	n := len(array)

	minVal, maxVal := array[0], array[0]
	for _, val := range array {
		if val < minVal {
			minVal = val
		} else if val > maxVal {
			maxVal = val
		}
	}

	// Optimal bucket count I tested
	bucketCount := n / 30
	if bucketCount == 0 {
        	bucketCount = 1
    	}
    	buckets := make([][]int, bucketCount)

	for _, val := range array {
		index := int(int(bucketCount-1) * (val - minVal) / (maxVal - minVal))
		buckets[index] = append(buckets[index], val)
	}

	result := make([]int, 0, n)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			InsertionSort(bucket)
			result = append(result, bucket...)
		}
	}

	return result
}
