package glist

// G stolen from StabbyCutyou/generics for teaching and the lulz
type G interface{}

type Node struct {
	g    G
	prev *Node
	next *Node
}

type Listy struct {
	headptr *Node
	tailptr *Node
	length  int
}
