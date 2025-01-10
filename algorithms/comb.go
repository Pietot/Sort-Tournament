package sort

func CombSort(array []int) []int {
    gap := len(array)
    shrink := 1.3
    sorted := false

    for !sorted {
        gap = int(float64(gap) / shrink)
        if gap <= 1 {
            gap = 1
            sorted = true
        }

        for i := 0; i+gap < len(array); i++ {
            if array[i] > array[i+gap] {
                array[i], array[i+gap] = array[i+gap], array[i]
                sorted = false
            }
        }
    }

    return array
}
