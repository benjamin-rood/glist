package glist

type List interface {
	Head() G
	Tail() G
	Append(G)
	AppendHead(G)
	Pop() G
	PopHead() G
	Length() int
}
