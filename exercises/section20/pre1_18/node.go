package pre1_18

type Any interface{}

type Node struct {
	Value Any
	Next  *Node
}

func New(value Any) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func NewWithNext(value Any, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}
