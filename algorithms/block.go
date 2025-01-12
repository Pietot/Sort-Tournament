package sort

func BlockSort(array []int) []int {
    blocks := splitIntoBlocks(array)
    for i, block := range blocks {
        blocks[i] = InsertionSort(block)
    }
    return mergeBlocks(blocks)
}

func splitIntoBlocks(array []int) [][]int {
    blockSize := 2
    blocks := make([][]int, 0)
    for i := 0; i < len(array); i += blockSize {
        end := i + blockSize
        if end > len(array) {
            end = len(array)
        }
        blocks = append(blocks, array[i:end])
    }
    return blocks
}

func mergeBlocks(blocks [][]int) []int {
    result := make([]int, 0)
    for _, block := range blocks {
        result = append(result, block...)
    }
    return result
}