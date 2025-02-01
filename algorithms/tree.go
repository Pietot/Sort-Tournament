package sort

func TreeSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	root := &Node{Value: array[0]}
	for _, val := range array[1:] {
		root.Insert(val)
	}

	result := make([]int, 0, len(array))
	root.Traverse(&result)
	return result
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if val < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) Traverse(result *[]int) {
	if n.Left != nil {
		n.Left.Traverse(result)
	}
	*result = append(*result, n.Value)
	if n.Right != nil {
		n.Right.Traverse(result)
	}
}
