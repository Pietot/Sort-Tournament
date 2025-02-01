package sort

func RedBlackTreeSort(array []int) []int {
    if len(array) == 0 {
        return array
    }

    tree := NewRBTree()
    for _, val := range array {
        tree.Insert(val)
    }

    result := make([]int, 0, len(array))
    tree.InOrderTraversal(func(val int) {
        result = append(result, val)
    })
    return result
}

type Color bool

const (
    Red   Color = true
    Black Color = false
)

type RBNode struct {
    Value  int
    Color  Color
    Left   *RBNode
    Right  *RBNode
    Parent *RBNode
}

type RBTree struct {
    Root *RBNode
}

func NewRBTree() *RBTree {
    return &RBTree{}
}

func (t *RBTree) Insert(val int) {
    newNode := &RBNode{Value: val, Color: Red}
    if t.Root == nil {
        t.Root = newNode
    } else {
        insertNode(t.Root, newNode)
    }
    t.fixInsert(newNode)
}

func insertNode(root, node *RBNode) {
    if node.Value < root.Value {
        if root.Left == nil {
            root.Left = node
            node.Parent = root
        } else {
            insertNode(root.Left, node)
        }
    } else {
        if root.Right == nil {
            root.Right = node
            node.Parent = root
        } else {
            insertNode(root.Right, node)
        }
    }
}

func (t *RBTree) fixInsert(node *RBNode) {
    for node != t.Root && node.Parent.Color == Red {
        if node.Parent == node.Parent.Parent.Left {
            uncle := node.Parent.Parent.Right
            if uncle != nil && uncle.Color == Red {
                node.Parent.Color = Black
                uncle.Color = Black
                node.Parent.Parent.Color = Red
                node = node.Parent.Parent
            } else {
                if node == node.Parent.Right {
                    node = node.Parent
                    t.rotateLeft(node)
                }
                node.Parent.Color = Black
                node.Parent.Parent.Color = Red
                t.rotateRight(node.Parent.Parent)
            }
        } else {
            uncle := node.Parent.Parent.Left
            if uncle != nil && uncle.Color == Red {
                node.Parent.Color = Black
                uncle.Color = Black
                node.Parent.Parent.Color = Red
                node = node.Parent.Parent
            } else {
                if node == node.Parent.Left {
                    node = node.Parent
                    t.rotateRight(node)
                }
                node.Parent.Color = Black
                node.Parent.Parent.Color = Red
                t.rotateLeft(node.Parent.Parent)
            }
        }
    }
    t.Root.Color = Black
}

func (t *RBTree) rotateLeft(node *RBNode) {
    right := node.Right
    node.Right = right.Left
    if right.Left != nil {
        right.Left.Parent = node
    }
    right.Parent = node.Parent
    if node.Parent == nil {
        t.Root = right
    } else if node == node.Parent.Left {
        node.Parent.Left = right
    } else {
        node.Parent.Right = right
    }
    right.Left = node
    node.Parent = right
}

func (t *RBTree) rotateRight(node *RBNode) {
    left := node.Left
    node.Left = left.Right
    if left.Right != nil {
        left.Right.Parent = node
    }
    left.Parent = node.Parent
    if node.Parent == nil {
        t.Root = left
    } else if node == node.Parent.Right {
        node.Parent.Right = left
    } else {
        node.Parent.Left = left
    }
    left.Right = node
    node.Parent = left
}

func (t *RBTree) InOrderTraversal(visit func(int)) {
    inOrderTraversal(t.Root, visit)
}

func inOrderTraversal(node *RBNode, visit func(int)) {
    if node != nil {
        inOrderTraversal(node.Left, visit)
        visit(node.Value)
        inOrderTraversal(node.Right, visit)
    }
}