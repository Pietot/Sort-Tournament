package sort

func CountingSortUnstable(array []int) []int {
    max := getMax(array)
    count := make([]int, max+1)
    output := make([]int, len(array))

    for _, val := range array {
        count[val]++
    }

    index := 0
    for i, c := range count {
        for c > 0 {
            output[index] = i
            index++
            c--
        }
    }

    return output
}
