package sort

func TreeSort(array []int) []int {
	root := &Node{Value: array[0]}
	for _, val := range array[1:] {
		root.Insert(val)
	}
	return root.Traverse()
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

func (n *Node) Traverse() []int {
	var result []int
	if n.Left != nil {
		result = append(result, n.Left.Traverse()...)
	}
	result = append(result, n.Value)
	if n.Right != nil {
		result = append(result, n.Right.Traverse()...)
	}
	return result
}
