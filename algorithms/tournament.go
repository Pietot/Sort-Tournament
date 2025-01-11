package sort

func TournamentSort(array []int) []int {
	n := len(array)
	tree := make([]int, 2*n-1)

	for i := 0; i < n; i++ {
		tree[n-1+i] = array[i]
	}
	for i := n - 2; i >= 0; i-- {
		tree[i] = min(tree[2*i+1], tree[2*i+2])
	}

	for i := 0; i < n; i++ {
		array[i] = tree[0]
		updateTree(tree, n)
	}

	return array
}

func updateTree(tree []int, n int) {
	index := 0
	for index < n-1 {
		left := 2*index + 1
		right := 2*index + 2
		if right < len(tree) && tree[right] < tree[left] {
			tree[index] = tree[right]
			index = right
		} else {
			tree[index] = tree[left]
			index = left
		}
	}
	tree[index] = int(^uint(0) >> 1)

	for index = (index - 1) / 2; index >= 0; index = (index - 1) / 2 {
		left := 2*index + 1
		right := 2*index + 2
		if right < len(tree) && tree[right] < tree[left] {
			tree[index] = tree[right]
		} else {
			tree[index] = tree[left]
		}
		if index == 0 {
			break
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
